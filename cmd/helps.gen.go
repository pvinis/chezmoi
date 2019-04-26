//go:generate go run github.com/twpayne/chezmoi/internal/extract-helps -i ../docs/REFERENCE.md -o helps.gen.go

package cmd

type help struct {
	long    string
	example string
}

var helps = map[string]help{
	"add": {
		long: "" +
			"Add targets to the source state. If any target is already in the source state,\n" +
			"then its source state is replaced with its current state in the destination\n" +
			"directory. The \"add\" command accepts additional flags:\n" +
			"\n" +
			"\"-e\", \"--empty\"\n" +
			"\n" +
			"Set the \"empty\" attribute on added files.\n" +
			"\n" +
			"\"-x\", \"--exact\"\n" +
			"\n" +
			"Set the \"exact\" attribute on added directories.\n" +
			"\n" +
			"\"-p\", \"--prompt\"\n" +
			"\n" +
			"Interactively prompt before adding each file.\n" +
			"\n" +
			"\"-r\", \"--recursive\"\n" +
			"\n" +
			"Recursively add all files, directories, and symlinks.\n" +
			"\n" +
			"\"-T\", \"--template\"\n" +
			"\n" +
			"Set the \"template\" attribute on added files and symlinks. In addition, \"chezmoi\"\n" +
			"attempts to automatically generate the template by replacing any template data\n" +
			"values with the equivalent template data keys. Longer subsitutions occur before\n" +
			"shorter ones.\n",
		example: "" +
			"  chezmoi add ~/.bashrc\n" +
			"  chezmoi add ~/.gitconfig --template\n" +
			"  chezmoi add ~/.vim --recursive\n" +
			"  chezmoi add ~/.oh-my-zsh --exact --recursive",
	},
	"apply": {
		long: "" +
			"Ensure that targets are in the target state, updating them if necessary. If no\n" +
			"targets are specified, the state of all targets are ensured.\n",
		example: "" +
			"  chezmoi apply\n" +
			"  chezmoi apply --dry-run --verbose\n" +
			"  chezmoi apply ~/.bashrc",
	},
	"archive": {
		long: "" +
			"Write a tar archive of the target state to stdout. This can be piped into \"tar\"\n" +
			"to inspect the target state.\n",
		example: "" +
			"  chezmoi archive | tar tvf -",
	},
	"cat": {
		long: "" +
			"Write the target state of targets  to stdout. targets must be files or symlinks.\n" +
			"For files, the target file contents are written. For symlinks, the target target\n" +
			"is written.\n",
		example: "" +
			"  chezmoi cat ~/.bashrc",
	},
	"cd": {
		long: "" +
			"Launch a shell in the source directory.\n",
		example: "" +
			"  chezmoi cd",
	},
	"chattr": {
		long: "" +
			"Change the attributes of targets. attributes specifies which attributes to\n" +
			"modify. Add attributes by specifying them or their abbreviations directly,\n" +
			"optionally prefixed with a plus sign (\"+\"). Remove attributes by prefixing them\n" +
			"or their attributes with the string \"no\" or a minus sign (\"-\"). The available\n" +
			"attributes and their abbreviations are:\n" +
			"\n" +
			"  Attribute    Abbreviation \n" +
			"  \"empty\"      \"e\"          \n" +
			"  \"encrypted\"  none         \n" +
			"  \"exact\"      none         \n" +
			"  \"executable\" \"x\"          \n" +
			"  \"private\"    \"p\"          \n" +
			"  \"template\"   \"t\"          \n" +
			"  \n" +
			"Multiple attributes modifications may be specified by separating them with a\n" +
			"comma (\",\").\n",
		example: "" +
			"  chezmoi chattr template ~/.bashrc\n" +
			"  chezmoi chattr noempty ~/.profile",
	},
	"completion": {
		long: "" +
			"Output shell completion code for the specified shell (\"bash\" or \"zsh\").\n",
		example: "" +
			"  chezmoi completion bash\n" +
			"  chezmoi completion zsh",
	},
	"data": {
		long: "" +
			"Write the computed template data in JSON format to stdout. The \"data\" command\n" +
			"accepts additional flags:\n" +
			"\n" +
			"\"-f\", \"--format\" format\n" +
			"\n" +
			"Print the computed template data in the given format. The accepted formats are\n" +
			"\"json\" (JSON), \"toml\" (TOML), and \"yaml\" (YAML).\n",
		example: "" +
			"  chezmoi data\n" +
			"  chezmoi data --format=yaml",
	},
	"diff": {
		long: "" +
			"Print the approximate shell commands required to ensure that targets in the\n" +
			"destination directory match the target state. If no targets are specifed,\n" +
			"print the commands required for all targets. It is equivalent to \"chezmoi apply\n" +
			"--dry-run --verbose\".\n",
		example: "" +
			"  chezmoi diff\n" +
			"  chezmoi diff ~/.bashrc",
	},
	"doctor": {
		long: "" +
			"Check for potential problems.\n",
		example: "" +
			"  chezmoi doctor",
	},
	"dump": {
		long: "" +
			"Dump the target state in JSON format. If no targets are specified, then the\n" +
			"entire target state. The \"dump\" command accepts additional arguments:\n" +
			"\n" +
			"\"-f\" / \"--format\" format\n" +
			"\n" +
			"Print the target state in the given format. The accepted formats are \"json\"\n" +
			"(JSON) and \"yaml\" (YAML).\n",
		example: "" +
			"  chezmoi dump ~/.bashrc\n" +
			"  chezmoi dump --format=yaml",
	},
	"edit": {
		long: "" +
			"Edit the source state of targets, which must be files or symlinks. The \"edit\"\n" +
			"command accepts additional arguments:\n" +
			"\n" +
			"\"-a\", \"--apply\"\n" +
			"\n" +
			"Apply target immediately after editing.\n" +
			"\n" +
			"\"-d\", \"--diff\"\n" +
			"\n" +
			"Print the difference between the target state and the actual state after\n" +
			"editing.\n" +
			"\n" +
			"\"-p\" / \"--prompt\"\n" +
			"\n" +
			"Prompt before applying each target.\n",
		example: "" +
			"  chezmoi edit ~/.bashrc\n" +
			"  chezmoi edit ~/.bashrc --apply --prompt",
	},
	"edit-config": {
		long: "" +
			"Edit the configuration file.\n",
		example: "" +
			"  chezmoi edit-config",
	},
	"forget": {
		long: "" +
			"Remove targets from the source state, i.e. stop managing them.\n",
		example: "" +
			"  chezmoi forget ~/.bashrc",
	},
	"help": {
		long: "" +
			"Print the help associated with command.\n",
	},
	"import": {
		long: "" +
			"FIXME document\n" +
			"\n" +
			"\"-x\" / \"--exact\"\n" +
			"\n" +
			"\"-r\", \"--remove-destination\"\n" +
			"\n" +
			"\"--strip-components\"\n",
	},
	"init": {
		long: "" +
			"Setup the source directory and update the destination directory to match the\n" +
			"target state. If repo is given then it is checked out into the source directory,\n" +
			"otherwise a new repository is initialized in the source directory. If a file\n" +
			"called \".chezmoi.format.tmpl\" exists, where \"format\" is one of the supported\n" +
			"file formats (e.g. \"json\", \"toml\", or \"yaml\") then a new configuration file is\n" +
			"created using that file as a template. Finally, if the \"--apply\" flag is passed,\n" +
			"\"chezmoi apply\" is run.\n",
		example: "" +
			"  chezmoi init https://github.com/user/dotfiles.git\n" +
			"  chezmoi init https://github.com/user/dotfiles.git --apply",
	},
	"merge": {
		long: "" +
			"Perform a three-way merge between the destination state, the source state, and\n" +
			"the target state. The merge tool is defined by the \"merge.command\" configuration\n" +
			"variable, and defaults to \"vimdiff\". If multiple targets are specified the merge\n" +
			"tool is invoked for each target. If the target state cannot be computed (for\n" +
			"example if source is a template containing errors or an encrypted file that\n" +
			"cannot be decrypted) a two-way merge is performed instead.\n",
		example: "" +
			"  chezmoi merge ~/.bashrc",
	},
	"remove": {
		long: "" +
			"Remove targets from both the source state and the destination directory.\n",
	},
	"secret": {
		long: "" +
			"Interact with a secret manager. See the \"Secret managers\" section for details.\n",
	},
	"source": {
		long: "" +
			"Execute the source version control system in the source directory with args.\n" +
			"Note that any flags for the source version control system must be sepeated with\n" +
			"a \"--\" to stop \"chezmoi\" from reading them.\n",
		example: "" +
			"  chezmoi source init\n" +
			"  chezmoi source add .\n" +
			"  chezmoi source commit -- -m \"Initial commit\"",
	},
	"source-path": {
		long: "" +
			"Print the path to each target's source state. If no targets are specified then\n" +
			"print the source directory.\n",
		example: "" +
			"  chezmoi source-path\n" +
			"  chezmoi source-path ~/.bashrc",
	},
	"unmanaged": {
		long: "" +
			"List all unmanaged files in the destination directory.\n",
		example: "" +
			"  chezmoi unmanaged",
	},
	"update": {
		long: "" +
			"Pull changes from the source VCS and apply any changes.\n",
		example: "" +
			"  chezmoi update",
	},
	"verify": {
		long: "" +
			"Verify that all targets match their target state. \"chezmoi\" exits with code 0\n" +
			"(success) if all targets match their target state, or 1 (failure) otherwise. If\n" +
			"no targets are specified then all targets are checked.\n",
		example: "" +
			"  chezmoi verify\n" +
			"  chezmoi verify ~/.bashrc",
	},
}
