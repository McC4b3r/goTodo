import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { V1Todo } from '../../../protos/clients/ts/src/models/V1Todo';

@Injectable({
  providedIn: 'root'
})
export class TodoService {
  private apiUrl = 'http://localhost:6001/v1/todos/list'

  constructor(private http: HttpClient) { }

  getTodos(): Observable<V1Todo[]> {
    return this.http.get<V1Todo[]>(this.apiUrl);
  }

  postTodo(todo: V1Todo): Observable<V1Todo> {
    return this.http.post<V1Todo>(this.apiUrl, todo);
  }

  deleteTodo(id: number): Observable<V1Todo> {
    return this.http.delete<V1Todo>(`${this.apiUrl}/${id}`);
  }

  updateTodo(todo: V1Todo): Observable<V1Todo> {
    return this.http.put<V1Todo>(`${this.apiUrl}/${todo.id}`, todo);
  }
}
