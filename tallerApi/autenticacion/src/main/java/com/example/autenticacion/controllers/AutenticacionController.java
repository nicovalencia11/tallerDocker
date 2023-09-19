package com.example.autenticacion.controllers;

import com.example.autenticacion.entities.Usuario;
import com.example.autenticacion.services.services.UsuarioService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/autenticacion")
public class AutenticacionController {

    @Autowired
    private UsuarioService usuarioService;

    /**
     * Metodo para realizar el login en el sistema
     * @param nombreUsuario
     * @param password
     * @return
     * @throws Exception
     */
    @PostMapping("/login")
    public ResponseEntity<String> login (@RequestBody String nombreUsuario, String password) throws Exception{
        usuarioService.loginUsuario(nombreUsuario, password);
        // retornar tokennnnn
        return ResponseEntity.status(HttpStatus.OK).body("Resgistro Exitoso");
    }

}
