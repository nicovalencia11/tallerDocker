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
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import java.util.Date;

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
    public ResponseEntity<String> login (@RequestBody String nombreUsuario, String password) throws Exception {
        Usuario usuario = usuarioService.loginUsuario(nombreUsuario, password);
        String token = AutenticacionController.generarTokenJWT(usuario);
        return ResponseEntity.status(HttpStatus.OK).body(token);
    }

    public static String generarTokenJWT(Usuario usuario) {
        String secretKey = "lUm27W{:{tVunMb6c()£WD1-rjl6H9Kkvci[+<?q"; // Use una clave de seguridad más segura en un entorno de producción.
        long tiempoDeExpiracion = 1000 * 60 * 60 * 10; // 10 horas

        return Jwts.builder()
                .setSubject(usuario.getNombre())
                .setIssuedAt(new Date(System.currentTimeMillis()))
                .setExpiration(new Date(System.currentTimeMillis() + tiempoDeExpiracion))
                .signWith(SignatureAlgorithm.HS512, secretKey)
                .compact();
    }

}
