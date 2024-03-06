package quare.software.boraracharbackend.presentation.controller

import client.borarachar.api.ScreensApi
import client.borarachar.model.*
import org.springframework.http.ResponseEntity
import org.springframework.stereotype.Controller

@Controller
class ScreensApiImpl : ScreensApi {
    override fun addCommentary(expenseId: String?, expenseCommentRequestBody: ExpenseCommentRequestBody?): ResponseEntity<Void> {
        return super.addCommentary(expenseId, expenseCommentRequestBody)
    }

    override fun creatExpenseScreen(userId: String?, expenseScreenRequestBody: ExpenseScreenRequestBody?): ResponseEntity<Void> {
        return super.creatExpenseScreen(userId, expenseScreenRequestBody)
    }

    override fun creatGroupScreen(userId: String?, groupsScreenRequestBody: GroupsScreenRequestBody?): ResponseEntity<Void> {
        return super.creatGroupScreen(userId, groupsScreenRequestBody)
    }

    override fun deleteExpenseScreen(expenseId: String?): ResponseEntity<Void> {
        return super.deleteExpenseScreen(expenseId)
    }

    override fun getGroupListScreen(userId: String?): ResponseEntity<GroupsScreenResponse> {
        return ResponseEntity.ok(null)
    }

    override fun getGroupScreen(page: Int?, limit: Int?, userId: String?, groupId: String?): ResponseEntity<GroupDetailsScreenResponse> {
        return super.getGroupScreen(page, limit, userId, groupId)
    }

    override fun screensActivitiesGet(page: Int?, limit: Int?, userId: String?): ResponseEntity<ActivitiesScreenResponse> {
        return super.screensActivitiesGet(page, limit, userId)
    }

    override fun screensContactsGet(userId: String?, accessCode: String?, refreshToken: String?): ResponseEntity<ContactsScreenResponse> {
        return super.screensContactsGet(userId, accessCode, refreshToken)
    }

    override fun screensContactsPost(userId: String?, accessCode: String?, refreshToken: String?, contactsScreenRequestBody: ContactsScreenRequestBody?): ResponseEntity<Void> {
        return super.screensContactsPost(userId, accessCode, refreshToken, contactsScreenRequestBody)
    }

    override fun screensLoginPost(loginRequestBody: LoginRequestBody?): ResponseEntity<InfoLoginResponse> {
        return super.screensLoginPost(loginRequestBody)
    }

    override fun updateExpenseScreen(expenseId: String?, expenseScreenRequestBody: ExpenseScreenRequestBody?): ResponseEntity<Void> {
        return super.updateExpenseScreen(expenseId, expenseScreenRequestBody)
    }
}