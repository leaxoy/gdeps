#gdeps


### HEAD
    This is a go package manage tool.    
install it
 
    go get -u -d github.com/guoziyan/gdeps

### How to use.

    Usage:
	    gdeps install <package> [--save]
	    gdeps remove <package> [--save]
    Arguments:
	    init    --generator a package.json file.
	    install --install packages define in the package.json file.
	    search  --search package and list packages find in godoc.org.
	    update  --update installed packages.
    Options:
	    -h|--help 		show this message.
	    --version  		show version of current.