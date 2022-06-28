package models

public data class CustomSubject(val name: String, public val isAdmin: Boolean) {
    public fun getIsAdmin(): Boolean {
        return this.isAdmin
    }
}
