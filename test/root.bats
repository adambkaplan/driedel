setup() {
    load 'test_helper/common-setup'
    _common_setup
}

@test "driedel command" {
    run driedel
    assert_output --partial 'Driedel is a command line that lets you run the Hanukah driedel game from an
interactive terminal'
}
