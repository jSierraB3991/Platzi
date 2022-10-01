package org.jsierra3991;

public class RepeatValues {

    public static String repeatString(String string, int times) {

        if(times < 0) {
            throw new IllegalArgumentException("negative times is not allowed");
        }
        String response = "";
        for (int i = 0; i < times; i++) {
            response += string;
        }
        return response;
    }
}
