test "X$(basename -- "$0")" = "Xsetup.sh" && \
  echo ERROR: This file should be "source'd", NOT run as a script. && exit 1
if [ -f ~/bin-bash/govars.sh ]; then
  source ~/bin-bash/govars.sh
else
  export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
  export PATH="$GOPATH/bin":$PATH
fi
