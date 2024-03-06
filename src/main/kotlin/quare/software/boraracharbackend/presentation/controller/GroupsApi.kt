package quare.software.boraracharbackend.presentation.controller

import client.borarachar.api.GroupsApi
import client.borarachar.model.GroupRequestBody
import client.borarachar.model.ParticipantRequestBody
import client.borarachar.model.UpdateParticipantRequestBody
import org.springframework.http.ResponseEntity
import org.springframework.stereotype.Controller

@Controller
class GroupsApi : GroupsApi {
    override fun addParticipant(groupId: String?, participantRequestBody: ParticipantRequestBody?): ResponseEntity<Void> {
        return super.addParticipant(groupId, participantRequestBody)
    }

    override fun deleteGroup(groupId: String?): ResponseEntity<Void> {
        return super.deleteGroup(groupId)
    }

    override fun removeParticipant(groupId: String?, userId: String?): ResponseEntity<Void> {
        return super.removeParticipant(groupId, userId)
    }

    override fun updateGroup(groupId: String?, groupRequestBody: GroupRequestBody?): ResponseEntity<Void> {
        return super.updateGroup(groupId, groupRequestBody)
    }

    override fun updateParticipant(groupId: String?, userId: String?, updateParticipantRequestBody: UpdateParticipantRequestBody?): ResponseEntity<Void> {
        return super.updateParticipant(groupId, userId, updateParticipantRequestBody)
    }
}