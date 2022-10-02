package org.jsierra3991;

import org.junit.Test;

import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.MatcherAssert.assertThat;


public class NumberUtilsArabicsToRomansTest {

    @Test
    public void return_I_to_number_is_1_o_grow_and_minus_to_4() {
        var  result = NumberUtils.arabicsToRomans(1);

        result = NumberUtils.arabicsToRomans(2);
        assertThat(result, is("II"));

        result = NumberUtils.arabicsToRomans(3);
        assertThat(result, is("III"));
    }

    @Test
    public void return_V_to_number_is_5_o_grow_and_minus_to_9() {
        var  result = NumberUtils.arabicsToRomans(5);
        assertThat(result, is("V"));

        result = NumberUtils.arabicsToRomans(6);
        assertThat(result, is("VI"));

        result = NumberUtils.arabicsToRomans(7);
        assertThat(result, is("VII"));

        result = NumberUtils.arabicsToRomans(8);
        assertThat(result, is("VIII"));
    }

    @Test
    public void return_X_to_number_is_5_o_grow_and_minus_to_50() {
        var  result = NumberUtils.arabicsToRomans(10);
        assertThat(result, is("X"));

        result = NumberUtils.arabicsToRomans(11);
        assertThat(result, is("XI"));

        result = NumberUtils.arabicsToRomans(15);
        assertThat(result, is("XV"));

        result = NumberUtils.arabicsToRomans(16);
        assertThat(result, is("XVI"));

        result = NumberUtils.arabicsToRomans(23);
        assertThat(result, is("XXIII"));

        result = NumberUtils.arabicsToRomans(32);
        assertThat(result, is("XXXII"));
    }

    @Test
    public void return_L_to_number_is_50_o_grow_and_minus_to_100() {
        var  result = NumberUtils.arabicsToRomans(50);
        assertThat(result, is("L"));

        result = NumberUtils.arabicsToRomans(51);
        assertThat(result, is("LI"));

        result = NumberUtils.arabicsToRomans(55);
        assertThat(result, is("LV"));

        result = NumberUtils.arabicsToRomans(56);
        assertThat(result, is("LVI"));

        result = NumberUtils.arabicsToRomans(60);
        assertThat(result, is("LX"));

        result = NumberUtils.arabicsToRomans(70);
        assertThat(result, is("LXX"));

        result = NumberUtils.arabicsToRomans(80);
        assertThat(result, is("LXXX"));

        result = NumberUtils.arabicsToRomans(85);
        assertThat(result, is("LXXXV"));

        result = NumberUtils.arabicsToRomans(86);
        assertThat(result, is("LXXXVI"));
    }

    @Test
    public void return_L_to_number_is_100_o_grow() {
        var  result = NumberUtils.arabicsToRomans(126);
        assertThat(result, is("CXXVI"));

        result = NumberUtils.arabicsToRomans(2507);
        assertThat(result, is("MMDVII"));
    }

    @Test
    public void return_romanic_compose_number() {
        var  result = NumberUtils.arabicsToRomans(4);
        assertThat(result, is("IV"));

        result = NumberUtils.arabicsToRomans(19);
        assertThat(result, is("XIX"));

        result = NumberUtils.arabicsToRomans(24);
        assertThat(result, is("XXIV"));

        result = NumberUtils.arabicsToRomans(40);
        assertThat(result, is("XL"));

        result = NumberUtils.arabicsToRomans(49);
        assertThat(result, is("XLIX"));

        result = NumberUtils.arabicsToRomans(90);
        assertThat(result, is("XC"));

        result = NumberUtils.arabicsToRomans(400);
        assertThat(result, is("CD"));

        result = NumberUtils.arabicsToRomans(900);
        assertThat(result, is("CM"));
    }
}