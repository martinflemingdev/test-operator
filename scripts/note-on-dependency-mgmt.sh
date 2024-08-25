# A note on dependency management 
# operator-sdk init generates a go.mod file to be used with Go modules. 
# The --repo=<path> flag is required when creating a project outside of $GOPATH/src, 
# as scaffolded files require a valid module path. 
# Ensure you activate module support by running 
export GO111MODULE=on 
# before using the SDK.

