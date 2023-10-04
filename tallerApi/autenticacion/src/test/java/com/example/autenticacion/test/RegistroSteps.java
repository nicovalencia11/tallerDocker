package com.example.autenticacion.test;


import io.cucumber.java.en.Given;
import io.cucumber.spring.CucumberContextConfiguration;
import org.springframework.boot.test.context.SpringBootTest;

@CucumberContextConfiguration
@SpringBootTest(classes = RegistroSteps.class)
public class RegistroSteps {

    @Given("debe estar el micro activo")
    public void validarEstadoServicio(){
        System.out.println("holaaaa");
    }

}
