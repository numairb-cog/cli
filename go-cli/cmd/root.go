package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "11.6.2"
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "npm",
	Short: "npm - a package manager for JavaScript",
	Long: `npm is a package manager for JavaScript.
It helps you install, share, and manage packages in your projects.`,
	Version: version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.npmrc)")
	rootCmd.PersistentFlags().Bool("global", false, "operate in global mode")
	rootCmd.PersistentFlags().Bool("json", false, "output JSON")
	rootCmd.PersistentFlags().String("loglevel", "notice", "log level (silent, error, warn, notice, http, timing, info, verbose, silly)")
	rootCmd.PersistentFlags().Bool("force", false, "force operation")
	rootCmd.PersistentFlags().Bool("dry-run", false, "do not actually perform actions")
	rootCmd.PersistentFlags().String("registry", "https://registry.npmjs.org", "npm registry URL")
	rootCmd.PersistentFlags().Bool("save", true, "save to dependencies")
	rootCmd.PersistentFlags().Bool("save-dev", false, "save to devDependencies")
	rootCmd.PersistentFlags().Bool("save-exact", false, "save exact version")
	rootCmd.PersistentFlags().Bool("ignore-scripts", false, "ignore scripts")
	rootCmd.PersistentFlags().StringSlice("workspace", []string{}, "workspace(s) to operate on")
	rootCmd.PersistentFlags().Bool("workspaces", false, "operate on all workspaces")

	rootCmd.AddCommand(accessCmd)
	rootCmd.AddCommand(adduserCmd)
	rootCmd.AddCommand(auditCmd)
	rootCmd.AddCommand(bugsCmd)
	rootCmd.AddCommand(cacheCmd)
	rootCmd.AddCommand(ciCmd)
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(dedupeCmd)
	rootCmd.AddCommand(deprecateCmd)
	rootCmd.AddCommand(diffCmd)
	rootCmd.AddCommand(distTagCmd)
	rootCmd.AddCommand(docsCmd)
	rootCmd.AddCommand(doctorCmd)
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(explainCmd)
	rootCmd.AddCommand(exploreCmd)
	rootCmd.AddCommand(findDupesCmd)
	rootCmd.AddCommand(fundCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(helpCmd)
	rootCmd.AddCommand(helpSearchCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(installCiTestCmd)
	rootCmd.AddCommand(installTestCmd)
	rootCmd.AddCommand(linkCmd)
	rootCmd.AddCommand(llCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(orgCmd)
	rootCmd.AddCommand(outdatedCmd)
	rootCmd.AddCommand(ownerCmd)
	rootCmd.AddCommand(packCmd)
	rootCmd.AddCommand(pingCmd)
	rootCmd.AddCommand(pkgCmd)
	rootCmd.AddCommand(prefixCmd)
	rootCmd.AddCommand(profileCmd)
	rootCmd.AddCommand(pruneCmd)
	rootCmd.AddCommand(publishCmd)
	rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(rebuildCmd)
	rootCmd.AddCommand(repoCmd)
	rootCmd.AddCommand(restartCmd)
	rootCmd.AddCommand(rootCmd2)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(sbomCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(shrinkwrapCmd)
	rootCmd.AddCommand(starCmd)
	rootCmd.AddCommand(starsCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(teamCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(tokenCmd)
	rootCmd.AddCommand(undeprecateCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(unpublishCmd)
	rootCmd.AddCommand(unstarCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(viewCmd)
	rootCmd.AddCommand(whoamiCmd)

	addAliases()
}

func initConfig() {
}

func addAliases() {
	rootCmd.AddCommand(&cobra.Command{
		Use:    "author",
		Hidden: true,
		Run:    ownerCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "home",
		Hidden: true,
		Run:    docsCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "issues",
		Hidden: true,
		Run:    bugsCmd.Run,
	})

	for _, alias := range []string{"info", "show"} {
		rootCmd.AddCommand(&cobra.Command{
			Use:    alias,
			Hidden: true,
			Run:    viewCmd.Run,
		})
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:    "find",
		Hidden: true,
		Run:    searchCmd.Run,
	})

	for _, alias := range []string{"add", "i", "in", "ins", "inst", "insta", "instal"} {
		rootCmd.AddCommand(&cobra.Command{
			Use:    alias,
			Hidden: true,
			Run:    installCmd.Run,
		})
	}

	for _, alias := range []string{"unlink", "remove", "rm", "r", "un"} {
		rootCmd.AddCommand(&cobra.Command{
			Use:    alias,
			Hidden: true,
			Run:    uninstallCmd.Run,
		})
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:    "rb",
		Hidden: true,
		Run:    rebuildCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "list",
		Hidden: true,
		Run:    lsCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "ln",
		Hidden: true,
		Run:    linkCmd.Run,
	})

	for _, alias := range []string{"create", "innit"} {
		rootCmd.AddCommand(&cobra.Command{
			Use:    alias,
			Hidden: true,
			Run:    initCmd.Run,
		})
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:    "it",
		Hidden: true,
		Run:    installTestCmd.Run,
	})

	for _, alias := range []string{"cit", "ic"} {
		rootCmd.AddCommand(&cobra.Command{
			Use:    alias,
			Hidden: true,
			Run:    installCiTestCmd.Run,
		})
	}

	for _, alias := range []string{"up", "upgrade", "udpate"} {
		rootCmd.AddCommand(&cobra.Command{
			Use:    alias,
			Hidden: true,
			Run:    updateCmd.Run,
		})
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:    "c",
		Hidden: true,
		Run:    configCmd.Run,
	})

	for _, alias := range []string{"s", "se"} {
		rootCmd.AddCommand(&cobra.Command{
			Use:    alias,
			Hidden: true,
			Run:    searchCmd.Run,
		})
	}

	for _, alias := range []string{"tst", "t"} {
		rootCmd.AddCommand(&cobra.Command{
			Use:    alias,
			Hidden: true,
			Run:    testCmd.Run,
		})
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:    "ddp",
		Hidden: true,
		Run:    dedupeCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "v",
		Hidden: true,
		Run:    viewCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "run-script",
		Hidden: true,
		Run:    runCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "clean-install",
		Hidden: true,
		Run:    ciCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "clean-install-test",
		Hidden: true,
		Run:    installCiTestCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "x",
		Hidden: true,
		Run:    execCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "why",
		Hidden: true,
		Run:    explainCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "la",
		Hidden: true,
		Run:    llCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "verison",
		Hidden: true,
		Run:    versionCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "hlep",
		Hidden: true,
		Run:    helpCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "dist-tags",
		Hidden: true,
		Run:    distTagCmd.Run,
	})

	for _, alias := range []string{"rum", "urn"} {
		rootCmd.AddCommand(&cobra.Command{
			Use:    alias,
			Hidden: true,
			Run:    runCmd.Run,
		})
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:    "ogr",
		Hidden: true,
		Run:    orgCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "add-user",
		Hidden: true,
		Run:    adduserCmd.Run,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:    "sit",
		Hidden: true,
		Run:    installCiTestCmd.Run,
	})

	for _, alias := range []string{"isnt", "isnta", "isntal", "isntall"} {
		rootCmd.AddCommand(&cobra.Command{
			Use:    alias,
			Hidden: true,
			Run:    installCmd.Run,
		})
	}

	for _, alias := range []string{"install-clean", "isntall-clean"} {
		rootCmd.AddCommand(&cobra.Command{
			Use:    alias,
			Hidden: true,
			Run:    ciCmd.Run,
		})
	}
}

func printNotImplemented(cmdName string) {
	fmt.Printf("npm %s: This command is not yet fully implemented in the Go version.\n", cmdName)
	fmt.Println("This is a rewrite of the npm CLI from Node.js to Go.")
	fmt.Printf("The '%s' command structure is in place but requires full implementation.\n", cmdName)
}
