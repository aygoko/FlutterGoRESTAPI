class User {
  final String id;
  final String login;
  final String email;

  User({
    required this.id,
    required this.login,
    required this.email,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'] as String,
      login: json['name'] as String,
      email: json['email'] as String,
    );
  }
}