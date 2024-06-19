
if(!window.ctf01d_tp_api) window.ctf01d_tp_api = {};

window.ctf01d_tp_api.games_list = function() {
    return $.ajax({
        url: '/api/v1/games',
        method: 'GET',
    });
}

window.ctf01d_tp_api.game_info = function(game_id) {
    return $.ajax({
        url: '/api/v1/games/' + game_id,
        method: 'GET',
    });
}

window.ctf01d_tp_api.game_create = function(game_data) {
    return $.ajax({
        url: '/api/v1/games',
        method: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(game_data),
    });
}

window.ctf01d_tp_api.auth_signin = function(auth_data) {
    return $.ajax({
        url: '/api/v1/auth/signin',
        method: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(auth_data),
      });
}

window.ctf01d_tp_api.auth_session = function (auth_data) {
    return $.ajax({
        url: '/api/v1/auth/session',
        method: 'GET',
        contentType: 'application/json',
        data: JSON.stringify(auth_data),
      });
}

window.ctf01d_tp_api.services_list = function() {
    return $.ajax({
        url: '/api/services',
        method: 'GET',
    });
}

window.ctf01d_tp_api.service_create = function(service_data) {
    return $.ajax({
        url: '/api/services',
        method: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(service_data),
    });
}

window.ctf01d_tp_api.service_update = function(service_id, service_data) {
    return $.ajax({
        url: '/api/services/' + service_id,
        method: 'PUT',
        contentType: 'application/json',
        data: JSON.stringify(service_data),
    });
}

window.ctf01d_tp_api.service_delete = function(service_id) {
    return $.ajax({
        url: '/api/services/' + service_id,
        method: 'DELETE',
    });
}

window.ctf01d_tp_api.service_info = function(service_id) {
    return $.ajax({
        url: '/api/services/' + service_id,
        method: 'GET',
    });
}
