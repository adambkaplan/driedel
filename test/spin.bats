setup() {
    load 'test_helper/common-setup'
    _common_setup
}

@test "spin the driedel" {
    run driedel spin
    assert_output --regexp '^You spun (נ Nun|ג Gimel|ה Hey|ש Shin)!$'
}

@test "check spin randomization" {
    # Check that we have a properly randomized driedel
    # For 15 spins of a driedel, there less than 1 in one billion chance that a letter is not spun.
    # 1/(4^15) if you want to be pedantic. This has higher confidence than a typical CI system can offer.
    local nunCount=0
    local gimelCount=0
    local heyCount=0
    local shinCount=0
    for (( i=0; i < 15; i++ )); do
        run driedel spin
        if [ "$output" = "You spun נ Nun!" ]; then
            $(( nunCount++ )) || true
        fi
        if [ "$output" = "You spun ג Gimel!" ]; then
            $(( gimelCount++ )) || true
        fi
        if [ "$output" = "You spun ה Hey!" ]; then
            $(( heyCount++ )) || true
        fi
        if [ "$output" = "You spun ש Shin!" ]; then
            $(( shinCount++ )) || true
        fi
    done
    echo "Checking that we got at least one nun"
    assert_not_equal ${nunCount} 0
    echo "Checking that we got at least one gimel"
    assert_not_equal ${gimelCount} 0
    echo "Checking that we got at least one hey"
    assert_not_equal ${heyCount} 0
    echo "Checking that we got at least one shin"
    assert_not_equal ${shinCount} 0
}
