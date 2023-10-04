package com.example.autenticacion.test;

import io.cucumber.junit.Cucumber;
import io.cucumber.junit.CucumberOptions;
import org.junit.runner.RunWith;


@RunWith(Cucumber.class)
@CucumberOptions(features = "src/test/java/com/example/autenticacion/test" +
        "/registro.feature", glue = "com.example.autenticacion.test.RegistroSteps")
public class TestRunner {
}
