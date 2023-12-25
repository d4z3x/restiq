package helper

import (
	"fmt"
	"os"
	"restiq/shell"
	"slices"

	"github.com/spf13/viper"
)

type config struct {
	ResticBin string `mapstructure:"restic_bin"`
	Token     map[string]string
	Repos     map[string]Repo
}

type Repo struct {
	KeyID         string `mapstructure:"key_id"`
	AppKey        string `mapstructure:"app_key"`
	Bucket        string
	User          string
	State         bool
	Location      string
	Type          string
	LimitUpload   string `mapstructure:"limit_upload"`
	LimitDownload string `mapstructure:"limit_download"`
	Compression   string
	MinAge        string `mapstructure:"min_age"` // --keep-within
	MaxAge        string `mapstructure:"max_age"` // --keep-within
	Dirs          []string
	MaxSnapshots  string `mapstructure:"max_snapshots"` // --keep-last
	PackSize      string // --pack-size
}

/*

{"message_type":"summary","files_new":1,"files_changed":66618,"files_unmodified":0,"dirs_new":0,"dirs_changed":2399,"dirs_unmodified":0,"data_blobs":488,"tree_blobs":2370,"data_added":908559139,"total_files_processed":66619,"total_bytes_processed":133008360576,"total_duration":3698.966863546,"snapshot_id":"19e847a2616632f1a12e9c72341944edf498c7d0ddaca6491f52e8cbd54c763d"}
{"message_type":"status","seconds_elapsed":63,"seconds_remaining":80,"percent_done":0.43935583979522497,"total_files":6230,"files_done":3701,"total_bytes":3551517562,"bytes_done":1560379981,"current_files":["/mnt/user/ScannedDocumentsNew/ETrade US and Morgan/Etrade address update.pdf","/mnt/user/ScannedDocumentsNew/ETrade US and Morgan/Restricted Stock 101_ Five Essentials of Restricted Stock.pdf"]}
*/

var C config
var configEnv string

func ReadViperConfig() {
	v := viper.New()
	v2 := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".") // optionally look for config in the working directory

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = v.Unmarshal(&C)

	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	v2.SetConfigName("cred")
	v2.SetConfigType("env")
	v2.AddConfigPath(".") // optionally look for config in the working directory
	v2.AutomaticEnv()

	if err = v2.ReadInConfig(); err != nil {
		fmt.Printf("unable to decode ENV into struct, %v", err)
	}
	if v2.Get("RESTIC_PASSWORD") == "" {
		fmt.Println("No password!")
		os.Exit(1)
	}
	C.Token = make(map[string]string)
	C.Token["PUSHOVER_TOKEN"] = v2.Get("PUSHOVER_TOKEN").(string)
	C.Token["RESTIC_PASSWORD"] = v2.Get("RESTIC_PASSWORD").(string)
}

// func ResticBackupAll() {
// 	for k, vv := range C.Repos {
// 		fmt.Printf("Repo %v has URL: %s:%s\n", k, vv.Type, vv.Location)
// 		fmt.Printf("Include dirs: %s\n", vv.Dirs)

// 		cmdStrRepo := []string{C.ResticBin, "-r", fmt.Sprintf("%s:%s", vv.Type, vv.Location)}
// 		cmdStr := cmdStrRepo

// 		if vv.LimitUpload != "" {
// 			cmdStr = append(cmdStr, "--limit-upload", vv.LimitUpload)
// 		}

// 		cmdStr = append(cmdStr,
// 			"--limit-upload", vv.LimitUpload,
// 			"--compression", vv.Compression)
// 		cmdStr = append(cmdStr,
// 			"forget",
// 			"--keep-last", vv.MaxSnapshots,
// 			"--prune", "--json")

// 		RunCmd(cmdStr, configEnv)
// 	}
// }

func AddParamsPost(repoName string) (params []string) {
	if C.Repos[repoName].MaxSnapshots != "" {
		params = append(params, "--keep-last", C.Repos[repoName].MaxSnapshots,
			"--prune")
	}
	return
}

func AddParams(repoName string) (params []string) {
	if C.Repos[repoName].PackSize != "" {
		params = append(params, "--pack-size", C.Repos[repoName].PackSize)
	}
	if C.Repos[repoName].LimitUpload != "" {
		params = append(params, "--limit-upload", C.Repos[repoName].LimitUpload)
	}
	if C.Repos[repoName].LimitDownload != "" {
		params = append(params, "--limit-download", C.Repos[repoName].LimitDownload)
	}
	params = append(params,
		"--compression", C.Repos[repoName].Compression,
		"--json")
	params = append(params, C.Repos[repoName].Dirs...)
	return
}

func ResticBackupAllNew() {
	for k, vv := range C.Repos {
		fmt.Printf("Repo %v has URL: %s:%s\n", k, vv.Type, vv.Location)
		ResticBackupRepoNew(findRepoURL(vv.Location))
		fmt.Printf("Include dirs: %s\n", vv.Dirs)

		// cmdStrRepo := []string{C.ResticBin, "-r", fmt.Sprintf("%s:%s", vv.Type, vv.Location)}
		// cmdStr := cmdStrRepo

		// if vv.LimitUpload != "" {
		// 	cmdStr = append(cmdStr, "--limit-upload", vv.LimitUpload)
		// }

		// cmdStr = append(cmdStr,
		// 	"--limit-upload", vv.LimitUpload,
		// 	"--compression", vv.Compression)
		// cmdStr = append(cmdStr,
		// 	"forget",
		// 	"--keep-last", vv.MaxSnapshots,
		// 	"--prune", "--json")

		// RunCmd(cmdStr, configEnv)
	}
}

func ResticPassword() string {
	return fmt.Sprintf("RESTIC_PASSWORD=%s", C.Token["RESTIC_PASSWORD"])
}

func PushoverPassword() string {
	return fmt.Sprintf("PUSHOVER_TOKEN=%s", C.Token["PUSHOVER_TOKEN"])
}

func ResticBackupRepoNew(repoName string) {
	fmt.Printf("Include dirs: %s\n", C.Repos[repoName].Dirs)

	shellCmd := shell.NewRestic(C.ResticBin)
	shellCmd.WithRepo(findRepoURL(repoName)).
		WithEnv(ResticPassword()).
		WithResticSubCmd("backup").WithArgs(AddParams(repoName)).
		RunWithJson()

	if C.Repos[repoName].MaxSnapshots != "" {
		shellCmd := shell.NewRestic(C.ResticBin)
		shellCmd.WithRepo(findRepoURL(repoName)).
			WithEnv(ResticPassword()).
			WithResticSubCmd("forget").WithArgs(AddParamsPost(repoName)).
			Run()
	}
}

func findRepoURL(repoName string) (repoURL string) {
	// Support local to disk backup
	if C.Repos[repoName].Type == "" || C.Repos[repoName].Type == "none" || C.Repos[repoName].Type == "local" {
		repoURL = C.Repos[repoName].Location
	} else {
		repoURL = fmt.Sprintf("%s:%s", C.Repos[repoName].Type, C.Repos[repoName].Location)
	}
	fmt.Printf("Repo %#v has URL: %s\n", repoName, repoURL)
	return repoURL
}

func ResticStatsRepo(repoName string) {
	shellCmd := shell.NewRestic(C.ResticBin).WithRepo(findRepoURL(repoName)).WithEnv(configEnv)
	shellCmd.WithResticSubCmd("stats").WithEnv(ResticPassword()).WithArgs(AddParams(repoName)).Run()
}

func ResticStatsAll() {
	// var repoURL string
	for _, vv := range C.Repos {
		ResticStatsRepo(findRepoURL(vv.Type + ":" + vv.Location))
	}
}

func ResticSnapshots(args []string) {
	shellCmd := shell.NewRestic(C.ResticBin).WithEnv(ResticPassword()).WithRepo(findRepoURL(args[0]))
	shellCmd.WithResticSubCmd("snapshots").Run()
}

func ResticLsSnapshot(args []string) {
	shellCmd := shell.NewRestic(C.ResticBin).WithRepo(findRepoURL(args[0])).WithEnv(configEnv)
	shellCmd.WithResticSubCmd("stats").WithResticOptionalArg(args).Run()
}

func ResticListRepos() {

	repoKeys := make([]string, 0, len(C.Repos))
	for k := range C.Repos {
		repoKeys = append(repoKeys, k)
	}

	slices.Sort(repoKeys)

	var repoURL string
	var count uint

	for _, v := range repoKeys {
		if C.Repos[v].Type == "" || C.Repos[v].Type == "none" || C.Repos[v].Type == "local" {
			repoURL = C.Repos[v].Location
		} else {
			repoURL = fmt.Sprintf("%s:%s", C.Repos[v].Type, C.Repos[v].Location)
		}

		fmt.Printf("%02d %-20s %s\n", count, v, repoURL)
		count++
	}
}
