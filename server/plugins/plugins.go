package plugins

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
	"github.com/stateprism/shell_vault/server/providers"
	"go.uber.org/zap"
	"os"
	"path"
	"slices"
)

type PluginExpr struct {
	File      string   `toml:"file"`
	AppliesTo []string `toml:"applies_to"`
}

type PluginManifest struct {
	Name        string       `toml:"name"`
	Version     string       `toml:"version"`
	Author      string       `toml:"author"`
	Description string       `toml:"description"`
	ExprSources []PluginExpr `toml:"sources"`
}

type Program struct {
	AppliesTo []string
	Prog      *vm.Program
}

type Plugin struct {
	Meta     PluginManifest
	Programs []Program
}

type Provider struct {
	config  providers.ConfigurationProvider
	plugins []*Plugin
	logger  *zap.Logger
}

type ProviderParams struct {
	Config providers.ConfigurationProvider
	Logger *zap.Logger
}

type Env struct {
	Method  string
	Session providers.SessionInfo
}

func NewProvider(p ProviderParams) (*Provider, error) {
	self := &Provider{
		config: p.Config,
		logger: p.Logger,
	}

	err := self.loadPlugins()
	if err != nil {
		return nil, err
	}

	return self, nil
}

func (p *Provider) Check(method string, env Env) (bool, error) {
	for _, p := range p.plugins {
		for _, prog := range p.Programs {
			if slices.Contains(prog.AppliesTo, method) {
				res, err := expr.Run(prog.Prog, env)
				if err != nil {
					return false, err
				}
				return res.(bool), nil
			}
		}
	}
	return false, nil
}

func (p *Provider) loadPlugins() error {
	pluginsPath := path.Join(p.config.GetStringOrDefault("paths.config", ""), "authorizers")

	if plDir, err := os.Stat(pluginsPath); err != nil || !plDir.IsDir() {
		return err
	}

	p.logger.Info(fmt.Sprintf("Looking for plugins at: %s", pluginsPath))
	dir, err := os.ReadDir(pluginsPath)
	if err != nil {
		return err
	}

	p.logger.Info(fmt.Sprintf("Found %d plugins to load at: %s", len(dir), pluginsPath))
	for _, entry := range dir {
		// skip single files
		if !entry.IsDir() {
			continue
		}
		plugin, err := p.loadPlugin(entry.Name())
		if err != nil {
			return err
		}
		p.plugins = append(p.plugins, plugin)
	}
	return nil
}

func (p *Provider) loadPlugin(pluginPath string) (*Plugin, error) {
	p.logger.Info("Trying to find manifest for plugin", zap.String("path", pluginPath))
	plugin := &Plugin{
		Programs: make([]Program, 0),
	}
	pluginsPath := path.Join(p.config.GetStringOrDefault("paths.config", ""), "authorizers")

	pluginFolder := path.Join(pluginsPath, pluginPath)
	data, err := os.ReadFile(path.Join(pluginFolder, "manifest.toml"))
	if err != nil {
		return nil, err
	}
	manifest := &PluginManifest{}

	_, err = toml.Decode(string(data), manifest)
	if err != nil {
		return nil, err
	}

	for _, source := range manifest.ExprSources {
		file, err := os.ReadFile(path.Join(pluginFolder, source.File))
		if err != nil {
			return nil, err
		}
		prog, err := expr.Compile(string(file), expr.Env(Env{}))
		if err != nil {
			p.logger.Error("Skipping loading of plugin", zap.String("name", manifest.Name), zap.Error(err))
			continue
		}
		plugin.Programs = append(plugin.Programs, Program{
			AppliesTo: source.AppliesTo,
			Prog:      prog,
		})
		p.logger.Info("Loaded plugin", zap.String("name", manifest.Name))
		p.logger.Debug("Manifest of last loaded plugin", zap.Any("manifest", manifest))
	}

	return plugin, nil
}
