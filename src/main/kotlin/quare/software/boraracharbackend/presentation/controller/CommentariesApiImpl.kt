package quare.software.boraracharbackend.presentation.controller

import client.borarachar.api.CommentariesApi
import client.borarachar.model.ExpenseCommentRequestBody
import org.springframework.http.ResponseEntity
import org.springframework.stereotype.Controller

@Controller
class CommentariesApiImpl: CommentariesApi {
    override fun deleteCommentary(commentaryId: String?): ResponseEntity<Void> {
        return super.deleteCommentary(commentaryId)
    }

    override fun updateCommentary(commentaryId: String?, expenseCommentRequestBody: ExpenseCommentRequestBody?): ResponseEntity<Void> {
        return ResponseEntity.ok(null)
    }
}