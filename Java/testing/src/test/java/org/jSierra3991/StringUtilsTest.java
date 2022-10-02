package org.jSierra3991;


import org.jsierra3991.StringUtils;
import org.junit.Assert;
import org.junit.Test;

public class StringUtilsTest {

    @Test
    public void test_string_one() {
        Assert.assertEquals("hola", StringUtils.repeatString("hola", 1));
    }

    @Test
    public void test_string_multiple_time() {
        Assert.assertEquals("holaholahola", StringUtils.repeatString("hola", 3));
    }

    @Test
    public void test_string_zero_times() {
        Assert.assertEquals("", StringUtils.repeatString("hola", 0));
    }

    @Test(expected = IllegalArgumentException.class)
    public void test_string_negative_times() {
        Assert.assertEquals("hola", StringUtils.repeatString("hola", -1));
    }
}
