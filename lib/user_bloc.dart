import 'dart:async';
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'user_model.dart';

enum UserEvent { fetchUsers }

class UserState {
  final List<User> users;
  final bool isLoading;
  final String? errorMessage;

  UserState({
    this.users = const [],
    this.isLoading = false,
    this.errorMessage,
  });

  UserState copyWith({
    List<User>? users,
    bool? isLoading,
    String? errorMessage,
  }) {
    return UserState(
      users: users ?? this.users,
      isLoading: isLoading ?? this.isLoading,
      errorMessage: errorMessage ?? this.errorMessage,
    );
  }
}

class UserBloc {
  final _stateStreamController = StreamController<UserState>();
  Stream<UserState> get stateStream => _stateStreamController.stream;
  final http.Client _client = http.Client();

  UserBloc() {
    _handleEvent(UserEvent.fetchUsers);
    _stateStreamController.add(UserState());
  }
  

  void refresh() {
    _handleEvent(UserEvent.fetchUsers);
  }


  void _handleEvent(UserEvent event) {
    switch (event) {
      case UserEvent.fetchUsers:
        _fetchUsers();
        break;
    }
  }

  void _fetchUsers() async {
  _stateStreamController.add(UserState(isLoading: true));

  try {
    final response = await http.get(Uri.parse('http://10.0.2.2:8080/api/users'));

    if (response.statusCode == 200) {
      final data = json.decode(response.body);
      if (data is List) { 
        final users = data.map((json) => User.fromJson(json)).toList();
        _stateStreamController.add(
          UserState(users: users, isLoading: false),
        );
      } else {
        _stateStreamController.add(
          UserState(
            isLoading: false,
            errorMessage: 'Invalid server response format',
          ),
        );
      }
    } else {
      _stateStreamController.add(
        UserState(
          isLoading: false,
          errorMessage: 'API Error: ${response.statusCode}',
        ),
      );
    }
  } catch (e) {
    _stateStreamController.add(
      UserState(
        isLoading: false,
        errorMessage: 'Fetch failed: $e',
      ),
    );
  }
}

  void dispose() {
    _stateStreamController.close();
  }
}