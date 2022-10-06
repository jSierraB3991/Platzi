package org.jsierra3991.functional.programming.proof;

import java.util.Arrays;
import java.util.List;
import java.util.function.BiFunction;
import java.util.function.Predicate;
import java.util.function.UnaryOperator;

public class StringFunctions {

    public static void run() {
        UnaryOperator<String> quote = text -> "\"" + text + "\"";
        System.out.println(quote.apply("Hello world"));

        BiFunction<Integer, Integer, Integer> multi = ( base, high ) -> base*high;
        System.out.println(multi.apply(5, 4));

        BiFunction<String, Integer, String> leftPad = (text, number) -> String.format("%" + number + "s", text);
        System.out.println(leftPad.apply("Java", 10));

        List<String> teachers = getLIst("Juan", "Nicolas", "Zulema");
        //Consumer<String> printer = text -> System.out.println(text);
        //teachers.forEach(printer);
        teachers.forEach(System.out::println);

        usingZeroArguments(() -> 2);
        //usingPredicate(text -> text.isEmpty());
        usingPredicate(String::isEmpty);
        usingBiFunction((base, high) -> base * high);
    }

    static <T> List<T> getLIst(T... elements) {
        return Arrays.asList(elements);
    }

    static void usingZeroArguments(ZeroArguments zeroArguments) { }
    static void usingPredicate(Predicate<String> predicate) {}
    static void usingBiFunction(BiFunction<Integer, Integer, Integer> biFunction) {}

    @FunctionalInterface
    interface ZeroArguments {
        int get();
    }
}