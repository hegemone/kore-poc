module Kore
  module Extension
    class IRCAdapter < Kore::Comm::PlatformAdapter
      include InjectPlatformClient[:irc_client]

      def name
        'irc'
      end

      def send_message(emsg)
        self.irc_client.send(self.serialize(emsg))
      end

      def listen
        self.log.debug "IRCAdapter#listen"
        self.irc_client.message_handler = method(:message_handler)
        self.irc_client.listen
      end

      private

      def serialize(emsg)
        "#{emsg.originator} - #{emsg.msg}"
      end

      def message_handler(raw_msg)
        # NOTE: Some amount of platform specific message processing...
        s = raw_msg.split(':')
        originator = s[0]
        msg = s[1].strip
        self.message_received(originator, msg)
      end
    end
  end
end
