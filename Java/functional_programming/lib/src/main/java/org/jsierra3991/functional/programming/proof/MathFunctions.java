package org.jsierra3991.functional.programming.proof;

import java.util.function.Function;
import java.util.function.Predicate;

public class MathFunctions {

    public static void run() {
        /*Functions */

        //OLd format
        Function<Integer, Integer> square = new Function<Integer, Integer>() {
            @Override
            public Integer apply(Integer x) {
                return x*x;
            }
        };
        //new format
        Function<Integer, Integer> squareFunction = (x) -> x * x;
        System.out.println(square.apply(5));
        System.out.println(squareFunction.apply(5));



        /* Predicate */
        Function<Integer, Boolean> isOdd = x -> x % 2 ==  0;
        Predicate<Integer> isEvent = x -> x%2 == 0;
        System.out.println(isOdd.apply(256));
        System.out.println(isEvent.test(256));

        Predicate<Student> approved = student -> student.getCalification() >= 6.0;
        Student sinue = new Student(5.9);
        System.out.println(approved.test(sinue));
        Student manuel = new Student(7.4);
        System.out.println(approved.test(manuel));
    }


    static class Student {
        private final double calification;

        public Student(double calification) {
            this.calification = calification;
        }

        public double getCalification() {
            return calification;
        }
    }
}