package com.example.autenticacion.entities;

import jakarta.persistence.*;
import lombok.*;

import java.io.Serializable;

@Entity
@Getter
@Setter
@ToString
@NoArgsConstructor
@AllArgsConstructor
@EqualsAndHashCode(onlyExplicitlyIncluded = true)
public class Usuario implements Serializable {

    @Id
    @EqualsAndHashCode.Include
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer codigo;

    @Column(nullable = false, unique = true, length = 45)
    private String nombre;

    @Column(nullable = false, unique = true, length = 250)
    private String correo;

    @Column(nullable = false, unique = true, length = 45)
    private String nombreUsuario;

    @Column(nullable = false)
    private String password;

}
