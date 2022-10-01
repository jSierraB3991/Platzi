package org.jSierra3991;


import org.jsierra3991.RepeatValues;
import org.junit.Assert;
import org.junit.Test;

public class RepeatValuesTest {

    @Test
    public void test_string_one() {
        Assert.assertEquals("hola", RepeatValues.repeatString("hola", 1));
    }

    @Test
    public void test_string_multiple_time() {
        Assert.assertEquals("holaholahola", RepeatValues.repeatString("hola", 3));
    }

    @Test
    public void test_string_zero_times() {
        Assert.assertEquals("", RepeatValues.repeatString("hola", 0));
    }

    @Test(expected = IllegalArgumentException.class)
    public void test_string_negative_times() {
        Assert.assertEquals("hola", RepeatValues.repeatString("hola", -1));
    }
}
