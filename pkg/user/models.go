package user

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

// id integer identificador único.
// name string Nombre del usuario, puede ser nulo.
// email string Email del usuario, se utiliza para iniciar sesión.
// password string Contraseña de usuario.
