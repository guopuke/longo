package version

var (
	gitTag       string = ""
	gitCommit    string = "$Format:%H$"          // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState string = "not a git tree"       // state of git tree, either "clean" or "dirty"
	buildData    string = "1970-01-01T00:00:00Z" // build date in IS080601 format, output of $(date -u + '%Y-%m-%d%H:%M:%SZ')
)
