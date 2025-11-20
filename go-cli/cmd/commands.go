package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var accessCmd = &cobra.Command{
	Use:   "access",
	Short: "Set access level on published packages",
	Long:  "Set access level on published packages (public or restricted)",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("access")
	},
}

var adduserCmd = &cobra.Command{
	Use:   "adduser",
	Short: "Add a registry user account",
	Long:  "Add a registry user account",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("adduser")
	},
}

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Run a security audit",
	Long:  "Run a security audit on your project dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("audit")
	},
}

var bugsCmd = &cobra.Command{
	Use:   "bugs",
	Short: "Report bugs for a package in a web browser",
	Long:  "Report bugs for a package in a web browser",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("bugs")
	},
}

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manipulates packages cache",
	Long:  "Manipulates packages cache",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("cache")
	},
}

var ciCmd = &cobra.Command{
	Use:   "ci",
	Short: "Clean install a project",
	Long:  "Clean install a project (similar to install but designed for CI environments)",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("ci")
	},
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Tab completion for npm",
	Long:  "Tab completion for npm",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("completion")
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage the npm configuration files",
	Long:  "Manage the npm configuration files",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("config")
	},
}

var dedupeCmd = &cobra.Command{
	Use:   "dedupe",
	Short: "Reduce duplication in the package tree",
	Long:  "Reduce duplication in the package tree",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("dedupe")
	},
}

var deprecateCmd = &cobra.Command{
	Use:   "deprecate",
	Short: "Deprecate a version of a package",
	Long:  "Deprecate a version of a package",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("deprecate")
	},
}

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "The registry diff command",
	Long:  "The registry diff command",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("diff")
	},
}

var distTagCmd = &cobra.Command{
	Use:   "dist-tag",
	Short: "Modify package distribution tags",
	Long:  "Modify package distribution tags",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("dist-tag")
	},
}

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Open documentation for a package in a web browser",
	Long:  "Open documentation for a package in a web browser",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("docs")
	},
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check your npm environment",
	Long:  "Check your npm environment",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("doctor")
	},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an installed package",
	Long:  "Edit an installed package",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("edit")
	},
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a command from a local or remote npm package",
	Long:  "Run a command from a local or remote npm package",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("exec")
	},
}

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain installed packages",
	Long:  "Explain installed packages",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("explain")
	},
}

var exploreCmd = &cobra.Command{
	Use:   "explore",
	Short: "Browse an installed package",
	Long:  "Browse an installed package",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("explore")
	},
}

var findDupesCmd = &cobra.Command{
	Use:   "find-dupes",
	Short: "Find duplication in the package tree",
	Long:  "Find duplication in the package tree",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("find-dupes")
	},
}

var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "Retrieve funding information",
	Long:  "Retrieve funding information",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("fund")
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value from the npm configuration",
	Long:  "Get a value from the npm configuration",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("get")
	},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Get help on npm",
	Long:  "Get help on npm",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			rootCmd.Help()
		} else {
			printNotImplemented("help")
		}
	},
}

var helpSearchCmd = &cobra.Command{
	Use:   "help-search",
	Short: "Search npm help documentation",
	Long:  "Search npm help documentation",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("help-search")
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a package.json file",
	Long:  "Create a package.json file",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("init")
	},
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a package",
	Long:  "Install a package and its dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("install")
	},
}

var installCiTestCmd = &cobra.Command{
	Use:   "install-ci-test",
	Short: "Install a project with a clean slate and run tests",
	Long:  "Install a project with a clean slate and run tests",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("install-ci-test")
	},
}

var installTestCmd = &cobra.Command{
	Use:   "install-test",
	Short: "Install package(s) and run tests",
	Long:  "Install package(s) and run tests",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("install-test")
	},
}

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Symlink a package folder",
	Long:  "Symlink a package folder",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("link")
	},
}

var llCmd = &cobra.Command{
	Use:   "ll",
	Short: "List installed packages with extended information",
	Long:  "List installed packages with extended information",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("ll")
	},
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to a registry user account",
	Long:  "Log in to a registry user account",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("login")
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of the registry",
	Long:  "Log out of the registry",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("logout")
	},
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List installed packages",
	Long:  "List installed packages",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("ls")
	},
}

var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "Manage orgs",
	Long:  "Manage orgs",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("org")
	},
}

var outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "Check for outdated packages",
	Long:  "Check for outdated packages",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("outdated")
	},
}

var ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "Manage package owners",
	Long:  "Manage package owners",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("owner")
	},
}

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Create a tarball from a package",
	Long:  "Create a tarball from a package",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("pack")
	},
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping npm registry",
	Long:  "Ping npm registry",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("ping")
	},
}

var pkgCmd = &cobra.Command{
	Use:   "pkg",
	Short: "Manages your package.json",
	Long:  "Manages your package.json",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("pkg")
	},
}

var prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Display prefix",
	Long:  "Display prefix",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("prefix")
	},
}

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Change settings on your registry profile",
	Long:  "Change settings on your registry profile",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("profile")
	},
}

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove extraneous packages",
	Long:  "Remove extraneous packages",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("prune")
	},
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a package",
	Long:  "Publish a package to the registry",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("publish")
	},
}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Retrieve a filtered list of packages",
	Long:  "Retrieve a filtered list of packages",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("query")
	},
}

var rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild a package",
	Long:  "Rebuild a package",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("rebuild")
	},
}

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Open package repository page in the browser",
	Long:  "Open package repository page in the browser",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("repo")
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart a package",
	Long:  "Restart a package",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("restart")
	},
}

var rootCmd2 = &cobra.Command{
	Use:   "root",
	Short: "Display npm root",
	Long:  "Display npm root",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("root")
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a package.json script",
	Long:  "Run a package.json script",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("run")
	},
}

var sbomCmd = &cobra.Command{
	Use:   "sbom",
	Short: "Generate a Software Bill of Materials (SBOM)",
	Long:  "Generate a Software Bill of Materials (SBOM)",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("sbom")
	},
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for packages",
	Long:  "Search for packages in the registry",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("search")
	},
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value in the npm configuration",
	Long:  "Set a value in the npm configuration",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("set")
	},
}

var shrinkwrapCmd = &cobra.Command{
	Use:   "shrinkwrap",
	Short: "Lock down dependency versions for publication",
	Long:  "Lock down dependency versions for publication",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("shrinkwrap")
	},
}

var starCmd = &cobra.Command{
	Use:   "star",
	Short: "Mark your favorite packages",
	Long:  "Mark your favorite packages",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("star")
	},
}

var starsCmd = &cobra.Command{
	Use:   "stars",
	Short: "View packages marked as favorites",
	Long:  "View packages marked as favorites",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("stars")
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a package",
	Long:  "Start a package",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("start")
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a package",
	Long:  "Stop a package",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("stop")
	},
}

var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Manage organization teams and team memberships",
	Long:  "Manage organization teams and team memberships",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("team")
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test a package",
	Long:  "Test a package",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("test")
	},
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Manage your authentication tokens",
	Long:  "Manage your authentication tokens",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("token")
	},
}

var undeprecateCmd = &cobra.Command{
	Use:   "undeprecate",
	Short: "Remove deprecation warning from a package",
	Long:  "Remove deprecation warning from a package",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("undeprecate")
	},
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove a package",
	Long:  "Remove a package",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("uninstall")
	},
}

var unpublishCmd = &cobra.Command{
	Use:   "unpublish",
	Short: "Remove a package from the registry",
	Long:  "Remove a package from the registry",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("unpublish")
	},
}

var unstarCmd = &cobra.Command{
	Use:   "unstar",
	Short: "Remove an item from your favorite packages",
	Long:  "Remove an item from your favorite packages",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("unstar")
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update packages",
	Long:  "Update packages",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("update")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Bump a package version",
	Long:  "Bump a package version",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("npm: %s\n", version)
			fmt.Printf("node: %s\n", "N/A (Go implementation)")
			fmt.Printf("go: %s\n", "1.21.5")
		} else {
			printNotImplemented("version")
		}
	},
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View registry info",
	Long:  "View registry info",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("view")
	},
}

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Display npm username",
	Long:  "Display npm username",
	Run: func(cmd *cobra.Command, args []string) {
		printNotImplemented("whoami")
	},
}
