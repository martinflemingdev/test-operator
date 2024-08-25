mkdir -p $HOME/projects/memcached-operator
cd $HOME/projects/memcached-operator
# we'll use a domain of example.com
# so all API groups will be <group>.example.com
operator-sdk init --domain mftest.com --repo github.com/martinflemingdev/test-operator