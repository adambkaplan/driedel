setup() {
    load 'test_helper/common-setup'
    _common_setup
}

@test "spin the driedel" {
    run driedel spin
    assert_output --regexp '^You spun (נ Nun|ג Gimel|ה Hey|ש Shin)!$'
}

