package com.example.autenticacion.controllers;

import com.example.autenticacion.entities.Usuario;
import com.example.autenticacion.services.services.UsuarioService;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
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

    @Autowired
    private RabbitTemplate rabbitTemplate;

    private final String exchangeName = "rootExchange";
    private final String routingKey = "tuRoutingKey";

    @PostMapping("/login")
    public ResponseEntity<String> login(@RequestBody Usuario usuarioRequest) throws Exception {
        try {
            Usuario usuario = usuarioService.loginUsuario(usuarioRequest.getNombreUsuario(), usuarioRequest.getPassword());
            String token = generarTokenJWT(usuario);
            usuario.setToken(token);
            usuarioService.actualizarUsuario(usuario);

            // Enviar mensaje de log exitoso a RabbitMQ
            String logMessage = "Login exitoso para el usuario: " + usuario.getNombreUsuario();
            rabbitTemplate.convertAndSend(exchangeName, routingKey, logMessage);

            return ResponseEntity.status(HttpStatus.OK).body(token);
        } catch (Exception e) {
            // Enviar mensaje de log fallido a RabbitMQ
            String logMessage = "Error en el login para el usuario: " + usuarioRequest.getNombreUsuario();
            rabbitTemplate.convertAndSend(exchangeName, routingKey, logMessage);

            throw e;
        }
    }

    public static String generarTokenJWT(Usuario usuario) {
        String secretKey = "lUm27WtVunMb"; // Asegúrate de usar una clave segura en producción.
        long tiempoDeExpiracion = 1000 * 60 * 60 * 10; // 10 horas
        return Jwts.builder()
                .setSubject(usuario.getNombreUsuario())
                .setIssuedAt(new Date(System.currentTimeMillis()))
                .setExpiration(new Date(System.currentTimeMillis() + tiempoDeExpiracion))
                .signWith(SignatureAlgorithm.HS512, secretKey)
                .compact();
    }
}
