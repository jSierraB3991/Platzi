package org.jsierra3991.functional.programming.proof;

import java.util.function.Consumer;
import java.util.function.Supplier;

public class CliArgumentsUtils {

    static void showHelp(CliArguments cliArgument) {
        Consumer<CliArguments> consumerHelper = cliArguments -> {
            if(cliArguments.isHelp()){
                System.out.println("Manual");
            }
        };
        consumerHelper.accept(cliArgument);
    }


    static CliArguments generateCli() {
        Supplier<CliArguments> generator = () -> new CliArguments();
        return generator.get();
    }
}