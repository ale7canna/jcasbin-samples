package models

public data class CustomSubject(val name: String, val isAdmin: Boolean, val age: Int, val address: String) {
    public fun getIsAdmin(): Boolean {
        return this.isAdmin
    }
}
