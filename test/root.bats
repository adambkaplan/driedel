setup() {
    load 'test_helper/bats-support/load'
    load 'test_helper/bats-assert/load'
    # get the containing directory of this file
    # use $BATS_TEST_FILENAME instead of ${BASH_SOURCE[0]} or $0,
    # as those will point to the bats executable's location or the preprocessed file respectively
    DIR="$( cd "$( dirname "$BATS_TEST_FILENAME" )" >/dev/null 2>&1 && pwd )"
    # make executables in _output/ visible to PATH
    PATH="$DIR/../_output:$PATH"
}

@test "driedel command" {
    run driedel
    assert_output --partial 'Driedel is a command line that lets you run the Hanukah driedel game from an
interactive terminal'
}
