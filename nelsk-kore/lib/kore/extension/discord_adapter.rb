module Kore
  module Extension
    class DiscordAdapter < Kore::Comm::PlatformAdapter
      include InjectPlatformClient[:discord_client]

      def name
        :discord
      end

      def send_message(emsg)
        self.discord_client.send(self.serialize(emsg))
      end

      def listen
        self.log.debug "DiscordAdapter#listen"
        #self.discord_client.message_handler = method(:message_handler)
        #self.discord_client.listen
      end

      private

      def serialize(emsg)
        "#{emsg.originator} - #{emsg.msg}"
      end

      def message_handler(raw_msg)
        # NOTE: Some amount of platform specific message processing...
        s = raw_msg.split(':')
        identity = s[0]
        msg = s[1].strip
        self.message_received(identity, msg)
      end
    end
  end
end
