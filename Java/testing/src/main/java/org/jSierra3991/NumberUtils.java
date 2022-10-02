package org.jSierra3991;

import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.Map;

public class NumberUtils {

    public static String arabicsToRomans(int number){
        HashMap<String, Integer> romanNumerals = new LinkedHashMap<>();
        romanNumerals.put("M", 1000);
        romanNumerals.put("CM", 900);
        romanNumerals.put("D", 500);
        romanNumerals.put("CD", 400);
        romanNumerals.put("C", 100);
        romanNumerals.put("XC", 90);
        romanNumerals.put("L", 50);
        romanNumerals.put("XL", 40);
        romanNumerals.put("X", 10);
        romanNumerals.put("IX", 9);
        romanNumerals.put("V", 5);
        romanNumerals.put("IV", 4);
        romanNumerals.put("I", 1);
        var result = "";

        for (Map.Entry<String, Integer> entry : romanNumerals.entrySet()) {
            String roman = entry.getKey();
            Integer arabic = entry.getValue();
            while (number >= arabic) {
                result += roman;
                number -= arabic;
            }
        }
        return result;
    }
}
