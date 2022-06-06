package main

type NotificationBuilder struct {
	title    string
	subtitle string
	message  string
	image    string
	icon     string
	priority int
	notType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (builder *NotificationBuilder) SetTitle(title string) {
	builder.title = title
}

func (builder *NotificationBuilder) SetSubtitle(subtitle string) {
	builder.subtitle = subtitle
}

func (builder *NotificationBuilder) SetMessage(message string) {
	builder.message = message
}

func (builder *NotificationBuilder) SetImage(image string) {
	builder.image = image
}

func (builder *NotificationBuilder) SetIcon(icon string) {
	builder.icon = icon
}

func (builder *NotificationBuilder) SetPriority(priority int) {
	builder.priority = priority
}

func (builder *NotificationBuilder) SetNotType(notType string) {
	builder.notType = notType
}

func (builder *NotificationBuilder) Build() (*Notification, error) {
	return &Notification{
		title:    builder.title,
		subtitle: builder.subtitle,
		message:  builder.message,
		image:    builder.image,
		icon:     builder.icon,
		priority: builder.priority,
		notType:  builder.notType,
	}, nil

}
