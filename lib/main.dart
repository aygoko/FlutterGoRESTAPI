import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'user_model.dart';
import 'user_form_screen.dart'; 
import 'user_list_screen.dart'; 
import 'package:flutter_keyboard_visibility/flutter_keyboard_visibility.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();

 
  const bool isImpellerDisabled = bool.fromEnvironment('FLUTTER_DISABLE_IMPELLER');
  if (!isImpellerDisabled) {
    debugPrint("Impeller is enabled. Consider disabling it if you encounter rendering issues.");
  }

 
  KeyboardVisibilityController().onChange.listen((visible) {
    debugPrint('Keyboard ${visible ? "opened" : "closed"}');
  });

  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'User Management',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: UserListScreen(),
    );
  }
}