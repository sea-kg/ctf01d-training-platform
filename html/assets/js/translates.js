window._tr = {
    "Игры": {"ru": "Игры", "en": "Games"},
    "Сервисы": {"ru": "Сервисы", "en": "Services"},
    "Команды": {"ru": "Команды", "en": "Teams"},
    "Войти": {"ru": "Войти", "en": "Sgin-in"},
    "Выйти": {"ru": "Выйти", "en": "Sign-out"},
    "Пользователи": {"ru": "Пользователи", "en": "Users"},
    "Мои команды": {"ru": "Мои команды", "en": "My Team(s)"},
    "Мои сервисы": {"ru": "Мои сервисы", "en": "My Service(s)"},
    "Что-то пошло не так...": {"ru": "Что-то пошло не так...", "en": "Something went wrong"},
    "Закрыть": {"ru": "Закрыть", "en": "Close"},
    "Результаты игры": {"ru": "Результаты игры", "en": "Game result"},
    "Добро пожаловать на тренировочную площадку ctf01d!": {
        "ru": "Добро пожаловать на тренировочную площадку ctf01d!",
        "en": "Welcome to ctf01d training platform!",
        "": ""
    },
    "Этот сервис может подготовить учебную игру 'атака-защита' на основе системы жюри ctf01d.": {
        "ru": "Этот сервис может подготовить учебную игру 'атака-защита' на основе системы жюри ctf01d.",
        "en": "This service can prepare training attack-defense game, based on ctf01d jury system",
        "ch": "该服务可以基于ctf01d陪审团系统准备训练攻防游戏",
        "by": "Гэты сэрвіс можа падрыхтаваць трэніровачную гульню ў атаку-абарону, заснаваную на сістэме журы ctf01d",
        "": ""
    },
    "Новая игра": {"ru": "Новая игра", "en": "New game"},
    "Такая страница не нашлась!": {
        "ru": "Такая страница не нашлась!",
        "en": "Page did not found!",
        "ch": "页面未找到！",
        "by": "Старонка не знойдзена!",
        "": ""
    },
    "Новый сервис": {"ru": "Новый сервис", "en": "New sevice"},
    "": {"ru": "", "en": ""}
};

function tr(cap) {
    var ret = window._tr[cap];
    if (ret === undefined) {
        console.warn("Not found translate for '" + cap + "'");
        return cap;
    }
    ret = ret['ru']; // TODO choose language
    return ret ? ret : cap;
}

function translateHtml() {
    var elems = document.querySelectorAll('[translate]');
    for (var i = 0; i < elems.length; i++) {
        var el = elems[i];
        var cap = el.getAttribute('translate');
        elems[i].innerHTML = tr(cap);
    }
}
