module Kore
  module Extension
    class BaconPlugin < Kore::Machinery::Base
      include Kore::Comm::Plugin

      match /bacon$/i,        :method => :command_bacon       # !bacon
      match /bacon\s+(\S+)/i, :method => :command_bacon_gift  # !bacon <user>

      def name
        :bacon
      end

      def help
        [
          help,
          'Usage: !bacon [user]'
        ].join "\n"
      end

      def command_bacon(msg, _)
        identity = msg.originator[:identity]
        self.send_message(msg.originator, "gives #{identity} a strip of delicious bacon.")
      end

      def command_bacon_gift(msg, match)
        to_user = match[1]
        identity = msg.originator[:identity]
        self.send_message(msg.originator, "gives #{to_user} a strip of delicious bacon as a gift from #{identity}.")
      end
    end
  end
end
