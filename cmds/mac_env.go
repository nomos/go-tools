package cmds

func init(){
	RegisterCmd("ibrew",setup_brew_mac)
}

var (
	setup_brew_mac = &WrappedCmd{
		CmdString: `
/usr/bin/expect << EOF
proc remote_exec {passwd} {
  spawn bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
  exp_internal 0
  expect {
    "*no':" { send "yes\n";exp_continue}
    "*password:" {send "${passwd}\n"}
    }
  close
}
remote_exec qq123123
EOF
`,
		Tips:      "ibrew [passwd]",
		ParamsNum: 1,
		ParamsMap: nil,
	}
)