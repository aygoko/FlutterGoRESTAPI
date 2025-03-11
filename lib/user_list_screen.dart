import 'package:flutter/material.dart';
import 'package:flutter_keyboard_visibility/flutter_keyboard_visibility.dart';
import 'user_bloc.dart';
import 'user_model.dart';
import 'user_form_screen.dart';

class UserListScreen extends StatefulWidget {
  @override
  _UserListScreenState createState() => _UserListScreenState();
}

class _UserListScreenState extends State<UserListScreen> {
  final UserBloc _userBloc = UserBloc();
  final KeyboardVisibilityController _keyboardVisibilityController =
      KeyboardVisibilityController();

  @override
  void initState() {
    super.initState();
    _keyboardVisibilityController.onChange.listen((visible) {
      print('Keyboard ${visible ? "opened" : "closed"}');
    });
  }

  @override
  void dispose() {
    _userBloc.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return StreamBuilder<UserState>(
      stream: _userBloc.stateStream,
      initialData: UserState(), 
      builder: (context, snapshot) {
        final state = snapshot.data ?? UserState();
        return Scaffold(
          appBar: AppBar(
            title: Text('Список пользователей'),
            actions: [
              IconButton(
                icon: Icon(Icons.refresh),
                onPressed: _userBloc.refresh,
              ),
            ],
          ),
          body: _buildContent(state),
          floatingActionButton: FloatingActionButton(
            child: Icon(Icons.add),
            onPressed: () async {
              final result = await Navigator.push(
                context,
                MaterialPageRoute(builder: (context) => UserFormScreen()),
              );
              if (result != null && result is Map && result['refresh'] == true) {
                _userBloc.refresh(); 
              }
            },
          ),
        );
      },
    );
  }

  Widget _buildContent(UserState state) {
    if (state.isLoading) {
      return Center(child: CircularProgressIndicator());
    }

    if (state.errorMessage != null) {
      return Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('Ошибка загрузки данных: ${state.errorMessage}'),
            SizedBox(height: 16),
            ElevatedButton(
              child: Text('Повторить'),
              onPressed: _userBloc.refresh,
            ),
          ],
        ),
      );
    }

    final users = state.users ?? []; 
    if (users.isEmpty) {
      return Center(child: Text('Нет данных'));
    }

    return Column(
      children: [
        Expanded(
          child: ListView.builder(
            itemCount: users.length,
            itemBuilder: (context, index) {
              final user = users[index];
              return Card(
                margin: EdgeInsets.all(8),
                child: ListTile(
                  title: Text(user.login),
                  subtitle: Text(user.email),
                  trailing: Text(user.id),
                ),
              );
            },
          ),
        ),
      ],
    );
  }
}