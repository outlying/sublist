package com.antyzero.sublist.app.domain

import java.util.UUID

sealed interface Item {
    data class Task(val description: String): Item
    data class TodoList(val id: UUID, val name: String, val items: List<Item>): Item
}

interface Repository {
    val masterList: Item.TodoList
    fun getAllTodoLists(): List<Item.TodoList>
    fun getTodoList(id: UUID): Item.TodoList?
    fun saveTodoList(todoList: Item.TodoList)
    fun deleteTodoList(id: UUID): Boolean
    fun addItemToList(todoListId: UUID, item: Item): Boolean
    fun removeItemFromList(todoListId: UUID, itemId: String): Boolean
}

class DevRepository : Repository {

    override val masterList = Item.TodoList(
        id = UUID.randomUUID(),
        name = "Master List",
        items = List(10) { index -> Item.Task("Task $index") }
    )

    private val todoLists = mutableListOf(masterList)

    override fun getAllTodoLists(): List<Item.TodoList> = todoLists

    override fun getTodoList(id: UUID): Item.TodoList? =
        todoLists.find { it.id == id }

    override fun saveTodoList(todoList: Item.TodoList) {
        val index = todoLists.indexOfFirst { it.id == todoList.id }
        if (index != -1) {
            todoLists[index] = todoList
        } else {
            todoLists.add(todoList)
        }
    }

    override fun deleteTodoList(id: UUID): Boolean {
        if (id == masterList.id) return false
        return todoLists.removeAll { it.id == id }
    }

    override fun addItemToList(todoListId: UUID, item: Item): Boolean {
        val todoList = getTodoList(todoListId) ?: return false
        val updatedList = todoList.copy(items = todoList.items + item)
        saveTodoList(updatedList)
        return true
    }

    override fun removeItemFromList(todoListId: UUID, itemId: String): Boolean {
        val todoList = getTodoList(todoListId) ?: return false
        val updatedItems = todoList.items.filterNot {
            it is Item.Task && it.description == itemId
        }
        if (updatedItems.size == todoList.items.size) return false
        saveTodoList(todoList.copy(items = updatedItems))
        return true
    }
}
