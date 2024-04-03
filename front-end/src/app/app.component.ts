import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { TodoListComponent } from './components/todo-list/todo-list.component';
import { TodoService } from './todo.service';
import { V1Todo } from '../../../protos/clients/ts/src/models';
import { catchError } from 'rxjs/operators';
import { of } from 'rxjs';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, TodoListComponent, FormsModule],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = `World's best Angular-Go Todo App`;
  todos: V1Todo[] | undefined;
  error: any;
  newTodoText: string = '';
  newTodo: V1Todo | undefined;

  constructor(private todoService: TodoService) { }

  ngOnInit(): void {
    this.fetchTodos();
  }

  fetchTodos(): void {
    this.todoService.getTodos().pipe(
      catchError((error) => {
        console.error('Error fetching todos', error);
        this.error = error;
        return of([]);
      })
    ).subscribe(todos => {
      console.log("todos:", todos)
      return (this.todos = todos);
    });
  }

  addTodo(): void {
    if (!this.newTodoText.trim()) {
      return; // Avoid adding empty todos
    }

    const newTodo: V1Todo = {
      todoName: this.newTodoText,
      todoType: 'TODO_GENERIC',
      priority: 1,
      completed: false,
    };

    if (!this.newTodoText.trim()) {
      return; // Avoid adding empty todos
    }

    this.todoService.addTodo(newTodo).subscribe({
      next: (todo) => {
        // console.log('Added new todo:', todo);
        console.log('this.todos:', this.todos);
        this.todos = this.todos ? [...this.todos, todo] as V1Todo[] : [todo] as V1Todo[];
        this.newTodoText = '';
      },
      error: (error) => {
        console.error('Error adding todo', error);
        this.error = error;
      }
    });
  }

}
