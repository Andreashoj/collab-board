import { writable, derived } from 'svelte/store';
import api from '$lib/api/api';

export interface Board {
	id: string;
	name: string;
	created_at: string;
	updated_at: string;
	members?: BoardMember[];
	logs?: BoardLog[];
}

export interface BoardMember {
	id: string;
	user_id: string;
	board_id: string;
	role: string;
	created_at: string;
	updated_at: string;
	user?: {
		id: string;
		email: string;
		firebase_uid: string;
	};
}

export interface BoardLog {
	id: string;
	board_id: string;
	user_id: string;
	change: string;
	created_at: string;
	user?: {
		id: string;
		email: string;
	};
}

interface BoardsState {
	boards: Board[];
	currentBoard: Board | null;
	currentMembers: BoardMember[];
	currentLogs: BoardLog[];
	loading: boolean;
	error: string | null;
}

const initialState: BoardsState = {
	boards: [],
	currentBoard: null,
	currentMembers: [],
	currentLogs: [],
	loading: true,
	error: null
};

function createBoardsStore() {
	const { subscribe, set, update } = writable<BoardsState>(initialState);

	return {
		subscribe,

		// Fetch all boards
		async fetchBoards() {
			update(state => ({ ...state, loading: true, error: null }));
			try {
				const response = await api.get('/boards');
				update(state => ({
					...state,
					boards: response.data || [],
					loading: false
				}));
			} catch (err: any) {
				console.error('Failed to fetch boards:', err);
				update(state => ({
					...state,
					error: err.response?.data?.error || 'Failed to load boards',
					loading: false
				}));
			}
		},

		// Fetch a single board with members and logs
		async fetchBoard(boardId: string) {
			update(state => ({ ...state, loading: true, error: null }));
			try {
				const [boardRes, membersRes, logsRes] = await Promise.all([
					api.get(`/boards/${boardId}`),
					api.get(`/board-members/board/${boardId}`).catch(() => ({ data: [] })),
					api.get(`/board-logs/board/${boardId}`).catch(() => ({ data: [] }))
				]);

				update(state => ({
					...state,
					currentBoard: boardRes.data,
					currentMembers: membersRes.data || [],
					currentLogs: logsRes.data || [],
					loading: false
				}));

				return boardRes.data;
			} catch (err: any) {
				console.error('Failed to fetch board:', err);
				const error = err.response?.status === 404 
					? 'Board not found'
					: err.response?.data?.error || 'Failed to load board';
				
				update(state => ({
					...state,
					error,
					loading: false
				}));
				throw err;
			}
		},

		// Create a new board
		async createBoard(name: string) {
			update(state => ({ ...state, loading: true, error: null }));
			try {
				const response = await api.post('/boards', { name });
				update(state => ({
					...state,
					boards: [...state.boards, response.data],
					loading: false
				}));
				return response.data;
			} catch (err: any) {
				console.error('Failed to create board:', err);
				update(state => ({
					...state,
					error: err.response?.data?.error || 'Failed to create board',
					loading: false
				}));
				throw err;
			}
		},

		// Update a board
		async updateBoard(boardId: string, name: string) {
			update(state => ({ ...state, loading: true, error: null }));
			try {
				const response = await api.put(`/boards/${boardId}`, { name });
				update(state => ({
					...state,
					boards: state.boards.map(b => b.id === boardId ? response.data : b),
					currentBoard: state.currentBoard?.id === boardId ? response.data : state.currentBoard,
					loading: false
				}));
				return response.data;
			} catch (err: any) {
				console.error('Failed to update board:', err);
				update(state => ({
					...state,
					error: err.response?.data?.error || 'Failed to update board',
					loading: false
				}));
				throw err;
			}
		},

		// Delete a board
		async deleteBoard(boardId: string) {
			update(state => ({ ...state, loading: true, error: null }));
			try {
				await api.delete(`/boards/${boardId}`);
				update(state => ({
					...state,
					boards: state.boards.filter(b => b.id !== boardId),
					currentBoard: state.currentBoard?.id === boardId ? null : state.currentBoard,
					loading: false
				}));
			} catch (err: any) {
				console.error('Failed to delete board:', err);
				update(state => ({
					...state,
					error: err.response?.data?.error || 'Failed to delete board',
					loading: false
				}));
				throw err;
			}
		},

		// Clear current board
		clearCurrentBoard() {
			update(state => ({
				...state,
				currentBoard: null,
				currentMembers: [],
				currentLogs: []
			}));
		},

		// Reset store
		reset() {
			set(initialState);
		}
	};
}

export const boardsStore = createBoardsStore();

// Derived stores for easier access
export const boards = derived(boardsStore, $store => $store.boards);
export const currentBoard = derived(boardsStore, $store => $store.currentBoard);
export const currentMembers = derived(boardsStore, $store => $store.currentMembers);
export const currentLogs = derived(boardsStore, $store => $store.currentLogs);
export const boardsLoading = derived(boardsStore, $store => $store.loading);
export const boardsError = derived(boardsStore, $store => $store.error);
