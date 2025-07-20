// Hàm tạo role mới
async function createRole() {
    const name = document.getElementById('addName').value;
    const slug = document.getElementById('addSlug').value;
    const description = document.getElementById('addDesc').value;

    try {
        const response = await fetch('/admin/roles', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'X-CSRF-Token': document.querySelector('meta[name="csrf-token"]').content
            },
            body: JSON.stringify({
                name: name,
                slug: slug,
                description: description
            })
        });

        if (response.ok) {
            window.location.reload();
        } else {
            const error = await response.json();
            alert(error.error || 'Failed to create role');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred');
    }
}

// Hàm hiển thị modal chỉnh sửa
function showEditRoleModal(id, name, slug, description) {
    document.getElementById('editID').value = id;
    document.getElementById('editName').value = name;
    document.getElementById('editSlug').value = slug;
    document.getElementById('editDesc').value = description || '';
    
    // Hiển thị modal
    const editModal = new bootstrap.Modal(document.getElementById('editModal'));
    editModal.show();
}

// Hàm cập nhật role
async function updateRole() {
    const id = document.getElementById('editID').value;
    const name = document.getElementById('editName').value;
    const slug = document.getElementById('editSlug').value;
    const description = document.getElementById('editDesc').value;

    try {
        const response = await fetch(`/admin/roles/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'X-CSRF-Token': document.querySelector('meta[name="csrf-token"]').content
            },
            body: JSON.stringify({
                name: name,
                slug: slug,
                description: description
            })
        });

        if (response.ok) {
            window.location.reload();
        } else {
            const error = await response.json();
            alert(error.error || 'Failed to update role');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred');
    }
}

// Hàm xóa role
async function deleteRole(id) {
    if (!confirm('Are you sure you want to delete this role?')) {
        return;
    }

    try {
        const response = await fetch(`/admin/roles/${id}`, {
            method: 'DELETE',
            headers: {
                'X-CSRF-Token': document.querySelector('meta[name="csrf-token"]').content
            }
        });

        if (response.ok) {
            window.location.reload();
        } else {
            const error = await response.json();
            alert(error.error || 'Failed to delete role');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred');
    }
}

// Hàm tìm kiếm role
async function searchRoles() {
    const query = document.getElementById('searchInput').value;
    const response = await fetch(`/admin/roles/search?q=${query}`);
    const data = await response.json();
    
    // Cập nhật bảng với kết quả tìm kiếm
    updateRoleTable(data);
}

function updateRoleTable(roles) {
    const tbody = document.getElementById('bulk-select-body');
    tbody.innerHTML = '';

    roles.forEach(role => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td class="align-middle white-space-nowrap">
                <div class="form-check mb-0">
                    <input class="form-check-input" type="checkbox" id="checkbox-${role.id}" />
                </div>
            </td>
            <td class="align-middle">${role.name}</td>
            <td class="align-middle">${role.slug}</td>
            <td class="align-middle">
                <button class="btn btn-outline-warning me-1 mb-1" type="button" 
                    onclick="showEditRoleModal('${role.id}', '${role.name}', '${role.slug}', '${role.description}')">
                    Edit
                </button>
                <button class="btn btn-outline-danger me-1 mb-1" type="button" 
                    onclick="deleteRole('${role.id}')">
                    Delete
                </button>
            </td>
        `;
        tbody.appendChild(row);
    });
}