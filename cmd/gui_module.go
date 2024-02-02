package cmd

import (
	"go.uber.org/fx"
	"protocol/assistant/internal/gui"
	"protocol/assistant/internal/gui/content"
	"protocol/assistant/internal/gui/dialog"
	"protocol/assistant/internal/gui/global_toolbar"
)

var uiModule = fx.Module("ui",
	fx.Provide(
		gui.NewUIApp,
		gui.NewUIWindow,
		gui.NewUIChan,
		gui.NewFxMenu(gui.NewFileMenu),
		gui.NewFxMenu(gui.NewHelpMenu),
		fx.Annotate(gui.NewMainMenu, fx.ParamTags(`group:"menus"`)),
		content.NewContentTab,
		fx.Annotate(content.NewMainContent, fx.ResultTags(`name:"mainContent"`)),
		fx.Annotate(global_toolbar.NewCreateConnectionAction,
			fx.ResultTags(`group:"toolbarActions"`),
		),
		dialog.NewCreateConnectionDialog,
		fx.Annotate(gui.NewMessageContent, fx.ResultTags(`name:"messageContent"`)),
		fx.Annotate(global_toolbar.NewToolBar,
			fx.ParamTags(`group:"toolbarActions"`),
			fx.ResultTags(`name:"toolBar"`),
		),
		fx.Annotate(gui.NewMainLayout, fx.ResultTags(`name:"mainLayout"`)),
		fx.Annotate(gui.NewStatusBar, fx.ResultTags(`name:"statusBar"`)),
		fx.Annotate(content.NewConnectionManager, fx.ResultTags(`name:"connectionManager"`)),
	),
	fx.Invoke(gui.SetupWindow, gui.ShowUI, gui.WindowShutdownHook),
)
