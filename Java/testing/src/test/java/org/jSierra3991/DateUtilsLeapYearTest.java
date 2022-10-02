package org.jSierra3991;

import org.junit.Test;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.CoreMatchers.is;

public class DateUtilsLeapYearTest {

    /*
        All years divisible by 400 ARE Leap years (1600, 2000, 2400)
        All years divisible by 100 but not by 400 are NOT leap years (1700, 1880, 1900),
        All years divisible by 4 but not by 100 ARE 1eap years (1996, 2004, 2008),
        All years not divisible by 4 are NOT Leap years (2017, 2018, 2019)
    */

    @Test
    // jupiter @DisplayName("return true when year is divisible by 400")
    public void return_true_when_year_is_divisible_by_400() {
        for (int i = 400; i <= 4000; i+=400) {
            var result = DateUtils.isLeapYear(i);
            assertThat(result, is(true));
        }
    }

    @Test
    //jupiter @DisplayName("return false when year is divisible by 100 but not for 400")
    public void return_false_when_year_is_divisible_by_100_but_not_for_400() {
        for (int i = 100; i <= 4000; i+=100) {
            if(i%400 != 0) {
                var result = DateUtils.isLeapYear(i);
                assertThat(result, is(false));
            }
        }
    }

    @Test
    //jupiter @DisplayName("return true when year is divisible by 4 but not for 100")
    public void return_true_when_year_is_divisible_by_4_but_not_for_100() {
        for (int i = 4; i <= 2022; i+=4) {
            if(i%100 != 0) {
                var result = DateUtils.isLeapYear(i);
                assertThat(result, is(true));
            }
        }
    }

    @Test
    //jupiter @DisplayName("return false when year not is divisible by 4")
    public void return_false_when_year_not_is_divisible_by_4() {
        for (int i = 1; i <= 2022; i++) {
            if(i%4 != 0) {
                var result = DateUtils.isLeapYear(i);
                assertThat(result, is(false));
            }
        }
    }
}