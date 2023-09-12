package com.example.autenticacion.controllers;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/usuario")
public class usuarioController {

    @GetMapping("/prueba")
    public ResponseEntity<String> prueba (){
        return ResponseEntity.status(HttpStatus.OK).body("prueba");
    }
}
