#!/bin/sh

##################################################
# Function to print the help for the fragmenter shell scripts
##################################################
function cHelp () {

  cat <<EOF
Invoke ". envsetup.sh" from your shell to add the following functions to your environment
- croot   Changes directory to the top of the tree
- godir      Go to the directory containing the specified file.

Helper Functions

- rn   Generate release notes and update the Changelog

- cHelp   Display this help message
EOF
}

##################################################
# Helper Functions
##################################################

#-------------------------------------------------
# Function to get to the top of the build tree
#-------------------------------------------------
function gettop() {
  git rev-parse --show-toplevel
}

#-------------------------------------------------
# Execute the specify function at the top of the tree
#-------------------------------------------------
function at_top() {

  T=$(gettop)

  local HERE=`pwd`

  # Go to the top of the build tree
  if [ -d "$T" ]; then

    cd $T

    echo $1
    bash -c $1

    cd $HERE
  else
    echo "Couldn't locate the top of the tree. Try setting TOP."
  fi

}

#--------------------------------------------------
# Function to go to the directory containing the
# specified file
#--------------------------------------------------
function godir() {

  # Check to make sure a correct argument was provided
  if [[ -z "$1" ]]; then
    echo "Usage: godir <regex>"
    return
  fi

  # Make the index if it doesn't exist
  T=$(gettop)
  if [[ ! -f $T/filelist ]]; then
    echo -n "Creating index..."
    (cd $T; find -E . -type f \( ! -path '*/.*' \) > filelist)
    echo " Done"
    echo ""
  fi

  # Get the list of files from the index
  local lines
  lines=($(grep "$1" $T/filelist | sed -e 's/\/[^/]*$//' | sort | uniq))
  if [[ ${#lines[@]} = 0 ]]; then
    echo "Not found"
    return
  fi

  # Create a menu based on the file list
  local pathname
  local choice

  if [[ ${#lines[@]} > 1 ]]; then
    while [[ -z "$pathname" ]]; do
      local index=1
      local line
      for line in ${lines[@]}; do
        printf "%6s %s\n" "[$index]" $line
        index=$((index + 1))
      done
      echo
      echo -n "Select one: "
      unset choice
      read choice
      if [[ $choice -gt ${#lines[@]} || $choice -lt 1 ]]; then
        echo "Invalid choice"
        continue
      fi

      pathname=${lines[$choice]}
    done
  else
    pathname=${lines[1]}
  fi

  cd $T/$pathname
}

# --------------------------------------------------
# Run the specified test suite.
# All configuration is expected to have happened by this point
# --------------------------------------------------
function release_notes() {

  local COMMIT_SHA_OF_LAST_RELEASE=$(git merge-base master $1)
  local COMMIT_SHA_OF_LAST_COMMIT_IN_CURRENT_RELEASE=$(git rev-list -n 1 HEAD)

  echo "=================================================="
  echo  Generating Release notes for Commits ${COMMIT_SHA_OF_LAST_RELEASE}..${COMMIT_SHA_OF_LAST_COMMIT_IN_CURRENT_RELEASE}
  echo "=================================================="

  changelog-gen \
    -repo terraform-provider-pureport \
    -owner pureport \
    -branch develop \
    -changelog .ci/changelog.tmpl \
    -releasenote .ci/release-note.tmpl \
    -no-note-label "changelog: no-release-note" \
    $COMMIT_SHA_OF_LAST_RELEASE $COMMIT_SHA_OF_LAST_COMMIT_IN_CURRENT_RELEASE
}
