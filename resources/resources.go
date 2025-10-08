package resources

import (
	"math/rand"
	"time"
)

type Notification struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var DailyNotifications = []Notification{
	{1, "☀️ Buổi sáng năng lượng!", "Hôm nay bạn chọn: Đạp xe hay đạp... deadline? 😆"},
	{2, "🚴‍♂️ Goter gọi kìa!", "Xe đã sẵn sàng, chỉ chờ người “đủ tỉnh” để lên đường!"},
	{3, "😜 Cà phê chưa tỉnh?", "Thử đạp 5 phút xem, bảo đảm tỉnh hơn espresso! ☕"},
	{4, "🌤️ Một ngày tuyệt vời để vận động!", "Hoặc ít nhất… để mở Goter lên nhìn thử thách 😅"},
	{5, "💪 Hôm nay không lười nữa nhé!", "Lười 1 ngày là bình thường, lười 2 ngày là… thói quen đó nha 😏"},
	{6, "🏁 Challenge Time!", "Bạn đã sẵn sàng chinh phục thử thách hôm nay chưa?"},
	{7, "🔥 Bắt đầu thôi!", "Cái khó của thử thách là... bắt đầu. Mà bạn thì giỏi nhất khoản đó rồi 😉"},
	{8, "🌈 Hít thở sâu…", "Và đạp thôi! Thành phố hôm nay đang đẹp lắm đấy 😍"},
	{9, "😆 Cảnh báo vui:", "Ai chưa hoàn thành thử thách sẽ bị Goter “ghim” trong tim 💔"},
	{10, "🕒 Thời gian trôi nhanh lắm!", "Còn chờ gì mà chưa check-in hôm nay? ⏰"},
	{11, "💥 Năng lượng đang chờ bạn!", "Mở Goter, bật mood, và cùng đi một vòng nào 🚴"},
	{12, "🚨 Challenge Alert!", "Một thử thách mới vừa xuất hiện – dám thử không? 😎"},
	{13, "🧠 Cảm giác lười đang tới gần...", "Chạy ngay trước khi nó bắt kịp bạn! 🏃‍♂️💨"},
	{14, "💬 Tin nóng:", "Người chăm nhất hôm nay đã hoàn thành thử thách từ sớm 😏 còn bạn thì sao?"},
	{15, "😜 Một chút vận động thôi mà!", "Goter không cần bạn hoàn hảo, chỉ cần bạn… chịu làm!"},
	{16, "🌇 Chiều nay rảnh không?", "Một vòng đạp quanh phố cho vui đi — không tính calories đâu 😁"},
	{17, "🎯 Check-in đi bạn ơi!", "Cứ mỗi lần check-in là một lần chiến thắng bản thân 💪"},
	{18, "🏆 Bạn là người chiến thắng!", "Nếu bạn mở Goter ngay bây giờ 😏"},
	{19, "💡 Thử thách không làm bạn mệt...", "Chỉ có “suy nghĩ có nên làm không” mới khiến bạn mệt thôi 🤣"},
	{20, "😂 Giữa muôn vàn cái deadline...", "Có một “thử thách” đang mỉm cười chờ bạn 💕"},
	{21, "🌅 Bình minh không đợi ai", "Nhưng Goter thì vẫn ở đây, đợi bạn bắt đầu! 😎"},
	{22, "🍃 Deep breath...", "Một ngày mới, một thử thách mới – và một bạn mới năng lượng hơn!"},
	{23, "😎 Hôm nay thử khác hôm qua đi!", "Không cần nhiều, chỉ cần… bớt lười hơn chút 😁"},
	{24, "🚴‍♀️ Chạy không?", "Nếu bạn không đi, thử thách sẽ đi mất đó 😜"},
	{25, "💭 Đừng nghĩ nữa!", "Bắt đầu luôn đi, 5 phút sau bạn sẽ cảm ơn chính mình! ❤️"},
	{26, "🔋 Nạp năng lượng thôi nào!", "Goter đang bật chế độ vui – bạn bật chưa? 😆"},
	{27, "🕶️ Bạn có biết?", "Người vui vẻ nhất hôm nay là người vừa hoàn thành thử thách 😁"},
	{28, "🧭 Lên đường thôi!", "Cứ đi, rồi bạn sẽ thấy điều thú vị trên đường đi 🛣️"},
	{29, "💥 Hôm nay là ngày của bạn!", "Không cần thắng ai – chỉ cần thắng chính mình 💪"},
	{30, "🎉 Kết thúc một tháng tuyệt vời!", "Thử thách tháng mới sắp tới, sẵn sàng chưa nào? 🚴‍♂️"},
}

func GetRandomNotification() Notification {
	today := time.Now()
	year, month, day := today.Date()
	// Tìm ngày cuối cùng của tháng hiện tại
	lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, today.Location()).Day()
	if day == lastDay {
		// Trả về item thứ 30 (index 29)
		return DailyNotifications[29]
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := r.Intn(len(DailyNotifications) - 1) // Chỉ random từ 0 đến 28
	return DailyNotifications[idx]
}
