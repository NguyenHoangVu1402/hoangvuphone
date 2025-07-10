function addRole() {
    const roleName = document.getElementById('role-name').value;
    const roleSlug = document.getElementById('role-slug').value;
    const roleDescription = document.getElementById('role-description').value;

    if (!roleName || !roleSlug) {
        alert('Name and Slug are required.');
        return;
    }

    const roleData = {
        name: roleName,
        slug: roleSlug,
        description: roleDescription
    };

    fetch('/roles', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(roleData)
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            alert('Role added successfully!');
            location.reload();
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('An error occurred while adding the role.');
    });
}

function showEditRoleModal(id, name, slug, description) {
    document.getElementById('role-id').value = id;
    document.getElementById('edit-role-name').value = name;
    document.getElementById('edit-role-slug').value = slug;
    document.getElementById('edit-role-description').value = description;
    $('#edit-role-modal').modal('show');
}

function editRole() {
    const roleId = document.getElementById('role-id').value;
    const roleName = document.getElementById('edit-role-name').value;
    const roleSlug = document.getElementById('edit-role-slug').value;
    const roleDescription = document.getElementById('edit-role-description').value;

    if (!roleName || !roleSlug) {
        alert('Name and Slug are required.');
        return;
    }

    const roleData = {
        name: roleName,
        slug: roleSlug,
        description: roleDescription
    };

    fetch(`/roles/${roleId}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(roleData)
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            alert('Role updated successfully!');
            location.reload();
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('An error occurred while updating the role.');
    });
}

function deleteRole(id) {
    if (confirm('Are you sure you want to delete this role?')) {
        fetch(`/roles/${id}`, {
            method: 'DELETE'
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                alert(data.error);
            } else {
                alert('Role deleted successfully!');
                location.reload();
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('An error occurred while deleting the role.');
        });
    }
}