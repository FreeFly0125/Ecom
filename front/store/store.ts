import { combineReducers, configureStore } from "@reduxjs/toolkit";
import storage from "../util/storage";
import userAuth from "./userAuth";
import { persistReducer, persistStore } from "redux-persist";
import getConfig from "next/config";
import products from "./products";
import cart from "./cart";
import toggleCart from "./toggleCart";
const { publicRuntimeConfig } = getConfig();

const rootReducer = combineReducers({
  userAuth,
  products,
  cart,
  toggleCart,
  // any other reducers here
});
const persistConfig = {
  key: "root",
  version: 1,
  storage,
};
const persistedReducer = persistReducer(persistConfig, rootReducer);

export const store = configureStore({
  reducer: persistedReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: ["persist/PERSIST"],
        ignoredPaths: [],
      },
    }),
  devTools: publicRuntimeConfig.appMode === "development",
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;
export let persistor = persistStore(store);
// export const wrapper = createWrapper(() => store);
