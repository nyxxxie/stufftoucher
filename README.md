# stufftoucher
After pwning a user during a pentest, it's often fun to dig through their shit
and try to find files and permissions that will enable further pwnage of other
things.  Stufftoucher is a tool that pokes around the system and does all that
digging for us (or tries to).  It looks for files with sensitive-looking things in it, application data and configs that contain user credentials, lulzy
permissions (passwordless sudo, access to other home directories, etc), and
more!  The tool also uses golang, which means we can pretty easily cross compile
and deploy it on whatever weird platform you manage to get a shell.
