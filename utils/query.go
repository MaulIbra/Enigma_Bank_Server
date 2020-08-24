package utils

const(
	INSERT_TRANSACTION = `INSERT INTO m_transaction values(?,?,?,?,?,?)`
	SELECT_TRANSACTION_BY_ID = `SELECT * FROM m_transaction where user_id = ?`
	SELECT_TRANSACTION = `SELECT * from m_transaction`
	SELECT_STOCK_MENU = `Select stok,menu_name from m_menu where menu_id=? and menu_active=1`
	INSERT_USER          = `INSERT INTO m_user values(?,?,?,?)`
	INSERT_PROFILE       = `insert into m_profile values(?,?,?,?,?);`
	SELECT_USER_BY_EMAIL = `SELECT * FROM m_user where username = ?`
	SELECT_USER_BY_ID = `SELECT * FROM m_user where user_id = ?`
	SELECT_USER          = `SELECT 
    m_profile.user_id,m_user.username,m_profile.nama_lengkap,m_profile.jenis_kelamin,m_profile.alamat 
    FROM M_USER INNER JOIN M_PROFILE 
    ON M_USER.USER_ID = M_PROFILE.USER_ID;`
	DELETE_USER_PROFILE = `DELETE FROM m_profile where user_id = ?`
	DELETE_USER         = `DELETE FROM m_user where user_id = ?`
	UPDATE_BALANCE_USER = `UPDATE m_user SET balance = ? WHERE user_id = ?`
)
