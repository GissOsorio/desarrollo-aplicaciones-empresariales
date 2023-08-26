// apiRoutes.js
const BASE_URL = 'http://localhost:8031';

const API_ROUTES = {
        getAllLists: `${BASE_URL}/lists`,
        getListByUserId: (userId) => `${BASE_URL}/lists/userid/${userId}`,
        getList: listId => `${BASE_URL}/lists/${listId}`,
        addList: `${BASE_URL}/lists`,
        updateListStatusById: listId => `${BASE_URL}/lists/status/${listId}`,
        getAllElements: `${BASE_URL}/elements`,
        getElementByListId: listId => `${BASE_URL}/elements/listid/${listId}`,
        getElementByListIdToDo: listId => `${BASE_URL}/elements-todo/listid/${listId}`,
        getElementByListIdDoing: listId => `${BASE_URL}/elements-doing/listid/${listId}`,
        getElementByListIdDone: listId => `${BASE_URL}/elements-done/listid/${listId}`,
        addElement: `${BASE_URL}/elements`,
        updateElementStatusById: elementId => `${BASE_URL}/elements/status/${elementId}`,
};
export default API_ROUTES